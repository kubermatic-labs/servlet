/*
Copyright 2024 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apiexport

import (
	"context"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"
	"go.uber.org/zap"

	"k8c.io/servlet/internal/controllerutil"
	predicateutil "k8c.io/servlet/internal/controllerutil/predicate"
	"k8c.io/servlet/internal/resources/reconciling"
	kdpservicesv1alpha1 "k8c.io/servlet/sdk/apis/services/v1alpha1"

	kcpdevv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/cluster"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/kontext"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	ControllerName = "servlet-apiexport"
)

type Reconciler struct {
	localClient    ctrlruntimeclient.Client
	platformClient ctrlruntimeclient.Client
	log            *zap.SugaredLogger
	recorder       record.EventRecorder
	lcName         logicalcluster.Name
	apiExportName  string
	servletName    string
	prFilter       labels.Selector
}

// Add creates a new controller and adds it to the given manager.
func Add(
	mgr manager.Manager,
	platformCluster cluster.Cluster,
	lcName logicalcluster.Name,
	log *zap.SugaredLogger,
	apiExportName string,
	servletName string,
	prFilter labels.Selector,
) error {
	reconciler := &Reconciler{
		localClient:    mgr.GetClient(),
		platformClient: platformCluster.GetClient(),
		lcName:         lcName,
		log:            log.Named(ControllerName),
		recorder:       mgr.GetEventRecorderFor(ControllerName),
		apiExportName:  apiExportName,
		servletName:    servletName,
		prFilter:       prFilter,
	}

	hasARS := predicate.NewPredicateFuncs(func(object ctrlruntimeclient.Object) bool {
		publishedResource, ok := object.(*kdpservicesv1alpha1.PublishedResource)
		if !ok {
			return false
		}

		return publishedResource.Status.ResourceSchemaName != ""
	})

	_, err := builder.ControllerManagedBy(mgr).
		Named(ControllerName).
		WithOptions(controller.Options{
			// we reconcile a single object in the KDP platform, no need for parallel workers
			MaxConcurrentReconciles: 1,
		}).
		// Watch for changes to APIExport on the platform side to start/restart the actual syncing controllers;
		// the cache is already restricted by a fieldSelector in the main.go to respect the RBC restrictions,
		// so there is no need here to add an additional filter.
		WatchesRawSource(source.Kind(platformCluster.GetCache(), &kcpdevv1alpha1.APIExport{}, controllerutil.EnqueueConst[*kcpdevv1alpha1.APIExport]("dummy"))).
		// Watch for changes to PublishedResources on the local service cluster
		Watches(&kdpservicesv1alpha1.PublishedResource{}, controllerutil.EnqueueConst[ctrlruntimeclient.Object]("dummy"), builder.WithPredicates(predicateutil.ByLabels(prFilter), hasARS)).
		Build(reconciler)
	return err
}

func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	r.log.Debug("Processing")
	return reconcile.Result{}, r.reconcile(ctx)
}

func (r *Reconciler) reconcile(ctx context.Context) error {
	// find all PublishedResources
	pubResources := &kdpservicesv1alpha1.PublishedResourceList{}
	if err := r.localClient.List(ctx, pubResources, &ctrlruntimeclient.ListOptions{
		LabelSelector: r.prFilter,
	}); err != nil {
		return fmt.Errorf("failed to list PublishedResources: %w", err)
	}

	// filter out those PRs that have not yet been processed into an ARS
	filteredPubResources := []kdpservicesv1alpha1.PublishedResource{}
	for i, pubResource := range pubResources.Items {
		if pubResource.Status.ResourceSchemaName != "" {
			filteredPubResources = append(filteredPubResources, pubResources.Items[i])
		}
	}

	// for each PR, we note down the created ARS and also the GVKs of related resources
	arsList := sets.New[string]()
	claimedResources := sets.New[string]()

	// PublishedResources use kinds, but the PermissionClaims use resource names (plural),
	// so we must translate accordingly
	mapper := r.platformClient.RESTMapper()

	for _, pubResource := range filteredPubResources {
		arsList.Insert(pubResource.Status.ResourceSchemaName)

		for _, rr := range pubResource.Spec.Related {
			resource, err := mapper.ResourceFor(schema.GroupVersionResource{
				Resource: rr.Kind,
			})
			if err != nil {
				return fmt.Errorf("unknown related resource kind %q: %w", rr.Kind, err)
			}

			claimedResources.Insert(resource.Resource)
		}
	}

	// Related resources (Secrets, ConfigMaps) are namespaced and so the Servlet will
	// always need to be able to see and manage namespaces.
	if claimedResources.Len() > 0 {
		claimedResources.Insert("namespaces")
	}

	if arsList.Len() == 0 {
		r.log.Debug("No ready PublishedResources available.")
		return nil
	}

	// reconcile an APIExport in the platform
	factories := []reconciling.NamedAPIExportReconcilerFactory{
		r.createAPIExportReconciler(arsList, claimedResources, r.servletName, r.apiExportName),
	}

	wsCtx := kontext.WithCluster(ctx, r.lcName)

	if err := reconciling.ReconcileAPIExports(wsCtx, factories, "", r.platformClient); err != nil {
		return fmt.Errorf("failed to reconcile APIExport: %w", err)
	}

	// try to get the virtual workspace URL of the APIExport;
	// TODO: This controller should watch the APIExport for changes
	// and then update
	// if err := wait.PollImmediate(100*time.Millisecond, 3*time.Second, func() (done bool, err error) {
	// 	apiExport := &kcpdevv1alpha1.APIExport{}
	// 	key := types.NamespacedName{Name: exportName}

	// 	if err := r.platformClient.Get(wsCtx, key, apiExport); ctrlruntimeclient.IgnoreNotFound(err) != nil {
	// 		return false, err
	// 	}

	// 	// NotFound (yet)
	// 	if apiExport.Name == "" {
	// 		return false, nil
	// 	}

	// 	// not ready
	// 	if len(apiExport.Status.VirtualWorkspaces) == 0 {
	// 		return false, nil
	// 	}

	// 	// do something with the URL...

	// 	return true, nil
	// }); err != nil {
	// 	return fmt.Errorf("failed to wait for virtual workspace to be ready: %w", err)
	// }

	return nil
}