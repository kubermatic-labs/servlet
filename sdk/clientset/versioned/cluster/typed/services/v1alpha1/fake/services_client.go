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
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	kcpservicesv1alpha1 "k8c.io/servlet/sdk/clientset/versioned/cluster/typed/services/v1alpha1"
	servicesv1alpha1 "k8c.io/servlet/sdk/clientset/versioned/typed/services/v1alpha1"
)

var _ kcpservicesv1alpha1.ServicesV1alpha1ClusterInterface = (*ServicesV1alpha1ClusterClient)(nil)

type ServicesV1alpha1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ServicesV1alpha1ClusterClient) Cluster(clusterPath logicalcluster.Path) servicesv1alpha1.ServicesV1alpha1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ServicesV1alpha1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ServicesV1alpha1ClusterClient) PublishedResources() kcpservicesv1alpha1.PublishedResourceClusterInterface {
	return &publishedResourcesClusterClient{Fake: c.Fake}
}

var _ servicesv1alpha1.ServicesV1alpha1Interface = (*ServicesV1alpha1Client)(nil)

type ServicesV1alpha1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ServicesV1alpha1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ServicesV1alpha1Client) PublishedResources() servicesv1alpha1.PublishedResourceInterface {
	return &publishedResourcesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}