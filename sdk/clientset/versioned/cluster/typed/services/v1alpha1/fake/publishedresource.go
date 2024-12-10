//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubermatic Kubernetes Platform contributors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kcp-dev/logicalcluster/v3"
	servicesv1alpha1 "k8c.io/servlet/sdk/apis/services/v1alpha1"
	applyconfigurationsservicesv1alpha1 "k8c.io/servlet/sdk/applyconfiguration/services/v1alpha1"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"

	servicesv1alpha1client "k8c.io/servlet/sdk/clientset/versioned/typed/services/v1alpha1"
)

var publishedResourcesResource = schema.GroupVersionResource{Group: "services.kdp.k8c.io", Version: "v1alpha1", Resource: "publishedresources"}
var publishedResourcesKind = schema.GroupVersionKind{Group: "services.kdp.k8c.io", Version: "v1alpha1", Kind: "PublishedResource"}

type publishedResourcesClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *publishedResourcesClusterClient) Cluster(clusterPath logicalcluster.Path) servicesv1alpha1client.PublishedResourceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &publishedResourcesClient{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of PublishedResources that match those selectors across all clusters.
func (c *publishedResourcesClusterClient) List(ctx context.Context, opts metav1.ListOptions) (*servicesv1alpha1.PublishedResourceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(publishedResourcesResource, publishedResourcesKind, logicalcluster.Wildcard, opts), &servicesv1alpha1.PublishedResourceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicesv1alpha1.PublishedResourceList{ListMeta: obj.(*servicesv1alpha1.PublishedResourceList).ListMeta}
	for _, item := range obj.(*servicesv1alpha1.PublishedResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested PublishedResources across all clusters.
func (c *publishedResourcesClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(publishedResourcesResource, logicalcluster.Wildcard, opts))
}

type publishedResourcesClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *publishedResourcesClient) Create(ctx context.Context, publishedResource *servicesv1alpha1.PublishedResource, opts metav1.CreateOptions) (*servicesv1alpha1.PublishedResource, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootCreateAction(publishedResourcesResource, c.ClusterPath, publishedResource), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}

func (c *publishedResourcesClient) Update(ctx context.Context, publishedResource *servicesv1alpha1.PublishedResource, opts metav1.UpdateOptions) (*servicesv1alpha1.PublishedResource, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateAction(publishedResourcesResource, c.ClusterPath, publishedResource), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}

func (c *publishedResourcesClient) UpdateStatus(ctx context.Context, publishedResource *servicesv1alpha1.PublishedResource, opts metav1.UpdateOptions) (*servicesv1alpha1.PublishedResource, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootUpdateSubresourceAction(publishedResourcesResource, c.ClusterPath, "status", publishedResource), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}

func (c *publishedResourcesClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewRootDeleteActionWithOptions(publishedResourcesResource, c.ClusterPath, name, opts), &servicesv1alpha1.PublishedResource{})
	return err
}

func (c *publishedResourcesClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewRootDeleteCollectionAction(publishedResourcesResource, c.ClusterPath, listOpts)

	_, err := c.Fake.Invokes(action, &servicesv1alpha1.PublishedResourceList{})
	return err
}

func (c *publishedResourcesClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*servicesv1alpha1.PublishedResource, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootGetAction(publishedResourcesResource, c.ClusterPath, name), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}

// List takes label and field selectors, and returns the list of PublishedResources that match those selectors.
func (c *publishedResourcesClient) List(ctx context.Context, opts metav1.ListOptions) (*servicesv1alpha1.PublishedResourceList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootListAction(publishedResourcesResource, publishedResourcesKind, c.ClusterPath, opts), &servicesv1alpha1.PublishedResourceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &servicesv1alpha1.PublishedResourceList{ListMeta: obj.(*servicesv1alpha1.PublishedResourceList).ListMeta}
	for _, item := range obj.(*servicesv1alpha1.PublishedResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *publishedResourcesClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewRootWatchAction(publishedResourcesResource, c.ClusterPath, opts))
}

func (c *publishedResourcesClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*servicesv1alpha1.PublishedResource, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(publishedResourcesResource, c.ClusterPath, name, pt, data, subresources...), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}

func (c *publishedResourcesClient) Apply(ctx context.Context, applyConfiguration *applyconfigurationsservicesv1alpha1.PublishedResourceApplyConfiguration, opts metav1.ApplyOptions) (*servicesv1alpha1.PublishedResource, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(publishedResourcesResource, c.ClusterPath, *name, types.ApplyPatchType, data), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}

func (c *publishedResourcesClient) ApplyStatus(ctx context.Context, applyConfiguration *applyconfigurationsservicesv1alpha1.PublishedResourceApplyConfiguration, opts metav1.ApplyOptions) (*servicesv1alpha1.PublishedResource, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewRootPatchSubresourceAction(publishedResourcesResource, c.ClusterPath, *name, types.ApplyPatchType, data, "status"), &servicesv1alpha1.PublishedResource{})
	if obj == nil {
		return nil, err
	}
	return obj.(*servicesv1alpha1.PublishedResource), err
}