//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishedResource) DeepCopyInto(out *PublishedResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishedResource.
func (in *PublishedResource) DeepCopy() *PublishedResource {
	if in == nil {
		return nil
	}
	out := new(PublishedResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublishedResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishedResourceList) DeepCopyInto(out *PublishedResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublishedResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishedResourceList.
func (in *PublishedResourceList) DeepCopy() *PublishedResourceList {
	if in == nil {
		return nil
	}
	out := new(PublishedResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublishedResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishedResourceSpec) DeepCopyInto(out *PublishedResourceSpec) {
	*out = *in
	out.Resource = in.Resource
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(ResourceFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.Naming != nil {
		in, out := &in.Naming, &out.Naming
		*out = new(ResourceNaming)
		**out = **in
	}
	if in.Projection != nil {
		in, out := &in.Projection, &out.Projection
		*out = new(ResourceProjection)
		(*in).DeepCopyInto(*out)
	}
	if in.Related != nil {
		in, out := &in.Related, &out.Related
		*out = make([]RelatedResourceSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishedResourceSpec.
func (in *PublishedResourceSpec) DeepCopy() *PublishedResourceSpec {
	if in == nil {
		return nil
	}
	out := new(PublishedResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishedResourceStatus) DeepCopyInto(out *PublishedResourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishedResourceStatus.
func (in *PublishedResourceStatus) DeepCopy() *PublishedResourceStatus {
	if in == nil {
		return nil
	}
	out := new(PublishedResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegexResourceLocator) DeepCopyInto(out *RegexResourceLocator) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegexResourceLocator.
func (in *RegexResourceLocator) DeepCopy() *RegexResourceLocator {
	if in == nil {
		return nil
	}
	out := new(RegexResourceLocator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelatedResourceReference) DeepCopyInto(out *RelatedResourceReference) {
	*out = *in
	in.Name.DeepCopyInto(&out.Name)
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(ResourceLocator)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelatedResourceReference.
func (in *RelatedResourceReference) DeepCopy() *RelatedResourceReference {
	if in == nil {
		return nil
	}
	out := new(RelatedResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelatedResourceSpec) DeepCopyInto(out *RelatedResourceSpec) {
	*out = *in
	in.Reference.DeepCopyInto(&out.Reference)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelatedResourceSpec.
func (in *RelatedResourceSpec) DeepCopy() *RelatedResourceSpec {
	if in == nil {
		return nil
	}
	out := new(RelatedResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDeleteMutation) DeepCopyInto(out *ResourceDeleteMutation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDeleteMutation.
func (in *ResourceDeleteMutation) DeepCopy() *ResourceDeleteMutation {
	if in == nil {
		return nil
	}
	out := new(ResourceDeleteMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceFilter) DeepCopyInto(out *ResourceFilter) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceFilter.
func (in *ResourceFilter) DeepCopy() *ResourceFilter {
	if in == nil {
		return nil
	}
	out := new(ResourceFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceLocator) DeepCopyInto(out *ResourceLocator) {
	*out = *in
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(RegexResourceLocator)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceLocator.
func (in *ResourceLocator) DeepCopy() *ResourceLocator {
	if in == nil {
		return nil
	}
	out := new(ResourceLocator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMutation) DeepCopyInto(out *ResourceMutation) {
	*out = *in
	if in.Rudi != nil {
		in, out := &in.Rudi, &out.Rudi
		*out = new(ResourceRudiMutation)
		**out = **in
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = new(ResourceDeleteMutation)
		**out = **in
	}
	if in.Regex != nil {
		in, out := &in.Regex, &out.Regex
		*out = new(ResourceRegexMutation)
		**out = **in
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(ResourceTemplateMutation)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMutation.
func (in *ResourceMutation) DeepCopy() *ResourceMutation {
	if in == nil {
		return nil
	}
	out := new(ResourceMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMutationSpec) DeepCopyInto(out *ResourceMutationSpec) {
	*out = *in
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = make([]ResourceMutation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = make([]ResourceMutation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMutationSpec.
func (in *ResourceMutationSpec) DeepCopy() *ResourceMutationSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceMutationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceNaming) DeepCopyInto(out *ResourceNaming) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceNaming.
func (in *ResourceNaming) DeepCopy() *ResourceNaming {
	if in == nil {
		return nil
	}
	out := new(ResourceNaming)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceProjection) DeepCopyInto(out *ResourceProjection) {
	*out = *in
	if in.ShortNames != nil {
		in, out := &in.ShortNames, &out.ShortNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceProjection.
func (in *ResourceProjection) DeepCopy() *ResourceProjection {
	if in == nil {
		return nil
	}
	out := new(ResourceProjection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRegexMutation) DeepCopyInto(out *ResourceRegexMutation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRegexMutation.
func (in *ResourceRegexMutation) DeepCopy() *ResourceRegexMutation {
	if in == nil {
		return nil
	}
	out := new(ResourceRegexMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRudiMutation) DeepCopyInto(out *ResourceRudiMutation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRudiMutation.
func (in *ResourceRudiMutation) DeepCopy() *ResourceRudiMutation {
	if in == nil {
		return nil
	}
	out := new(ResourceRudiMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceTemplateMutation) DeepCopyInto(out *ResourceTemplateMutation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceTemplateMutation.
func (in *ResourceTemplateMutation) DeepCopy() *ResourceTemplateMutation {
	if in == nil {
		return nil
	}
	out := new(ResourceTemplateMutation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceResourceDescriptor) DeepCopyInto(out *SourceResourceDescriptor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceResourceDescriptor.
func (in *SourceResourceDescriptor) DeepCopy() *SourceResourceDescriptor {
	if in == nil {
		return nil
	}
	out := new(SourceResourceDescriptor)
	in.DeepCopyInto(out)
	return out
}