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

package projection

import (
	kdpservicesv1alpha1 "k8c.io/servlet/sdk/apis/services/v1alpha1"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// PublishedResourceSourceGVK returns the source GVK of the local resources
// that are supposed to be published.
func PublishedResourceSourceGVK(pubRes *kdpservicesv1alpha1.PublishedResource) schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   pubRes.Spec.Resource.APIGroup,
		Version: pubRes.Spec.Resource.Version,
		Kind:    pubRes.Spec.Resource.Kind,
	}
}

// PublishedResourceProjectedGVK returns the effective GVK after the projection
// rules have been applied according to the PublishedResource.
func PublishedResourceProjectedGVK(pubRes *kdpservicesv1alpha1.PublishedResource, platformAPIGroup string) schema.GroupVersionKind {
	apiVersion := pubRes.Spec.Resource.Version
	kind := pubRes.Spec.Resource.Kind

	if projection := pubRes.Spec.Projection; projection != nil {
		if v := projection.Version; v != "" {
			apiVersion = v
		}

		if k := projection.Kind; k != "" {
			kind = k
		}
	}

	return schema.GroupVersionKind{
		Group:   platformAPIGroup,
		Version: apiVersion,
		Kind:    kind,
	}
}
