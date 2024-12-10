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

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	servicesv1alpha1 "k8c.io/servlet/sdk/apis/services/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	servicesv1alpha1client "k8c.io/servlet/sdk/clientset/versioned/typed/services/v1alpha1"
)

// PublishedResourcesClusterGetter has a method to return a PublishedResourceClusterInterface.
// A group's cluster client should implement this interface.
type PublishedResourcesClusterGetter interface {
	PublishedResources() PublishedResourceClusterInterface
}

// PublishedResourceClusterInterface can operate on PublishedResources across all clusters,
// or scope down to one cluster and return a servicesv1alpha1client.PublishedResourceInterface.
type PublishedResourceClusterInterface interface {
	Cluster(logicalcluster.Path) servicesv1alpha1client.PublishedResourceInterface
	List(ctx context.Context, opts metav1.ListOptions) (*servicesv1alpha1.PublishedResourceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type publishedResourcesClusterInterface struct {
	clientCache kcpclient.Cache[*servicesv1alpha1client.ServicesV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *publishedResourcesClusterInterface) Cluster(clusterPath logicalcluster.Path) servicesv1alpha1client.PublishedResourceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).PublishedResources()
}

// List returns the entire collection of all PublishedResources across all clusters.
func (c *publishedResourcesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*servicesv1alpha1.PublishedResourceList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PublishedResources().List(ctx, opts)
}

// Watch begins to watch all PublishedResources across all clusters.
func (c *publishedResourcesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PublishedResources().Watch(ctx, opts)
}