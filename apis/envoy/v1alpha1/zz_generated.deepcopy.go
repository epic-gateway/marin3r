// +build !ignore_autogenerated

/*


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
	"github.com/operator-framework/operator-lib/status"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigRevisionRef) DeepCopyInto(out *ConfigRevisionRef) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigRevisionRef.
func (in *ConfigRevisionRef) DeepCopy() *ConfigRevisionRef {
	if in == nil {
		return nil
	}
	out := new(ConfigRevisionRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyBootstrap) DeepCopyInto(out *EnvoyBootstrap) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyBootstrap.
func (in *EnvoyBootstrap) DeepCopy() *EnvoyBootstrap {
	if in == nil {
		return nil
	}
	out := new(EnvoyBootstrap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyBootstrap) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyBootstrapList) DeepCopyInto(out *EnvoyBootstrapList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvoyBootstrap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyBootstrapList.
func (in *EnvoyBootstrapList) DeepCopy() *EnvoyBootstrapList {
	if in == nil {
		return nil
	}
	out := new(EnvoyBootstrapList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyBootstrapList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyBootstrapSpec) DeepCopyInto(out *EnvoyBootstrapSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyBootstrapSpec.
func (in *EnvoyBootstrapSpec) DeepCopy() *EnvoyBootstrapSpec {
	if in == nil {
		return nil
	}
	out := new(EnvoyBootstrapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyBootstrapStatus) DeepCopyInto(out *EnvoyBootstrapStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyBootstrapStatus.
func (in *EnvoyBootstrapStatus) DeepCopy() *EnvoyBootstrapStatus {
	if in == nil {
		return nil
	}
	out := new(EnvoyBootstrapStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfig) DeepCopyInto(out *EnvoyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfig.
func (in *EnvoyConfig) DeepCopy() *EnvoyConfig {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigList) DeepCopyInto(out *EnvoyConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvoyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigList.
func (in *EnvoyConfigList) DeepCopy() *EnvoyConfigList {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigRevision) DeepCopyInto(out *EnvoyConfigRevision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigRevision.
func (in *EnvoyConfigRevision) DeepCopy() *EnvoyConfigRevision {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigRevision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyConfigRevision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigRevisionList) DeepCopyInto(out *EnvoyConfigRevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvoyConfigRevision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigRevisionList.
func (in *EnvoyConfigRevisionList) DeepCopy() *EnvoyConfigRevisionList {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigRevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvoyConfigRevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigRevisionSpec) DeepCopyInto(out *EnvoyConfigRevisionSpec) {
	*out = *in
	if in.EnvoyAPI != nil {
		in, out := &in.EnvoyAPI, &out.EnvoyAPI
		*out = new(string)
		**out = **in
	}
	if in.Serialization != nil {
		in, out := &in.Serialization, &out.Serialization
		*out = new(string)
		**out = **in
	}
	if in.EnvoyResources != nil {
		in, out := &in.EnvoyResources, &out.EnvoyResources
		*out = new(EnvoyResources)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigRevisionSpec.
func (in *EnvoyConfigRevisionSpec) DeepCopy() *EnvoyConfigRevisionSpec {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigRevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigRevisionStatus) DeepCopyInto(out *EnvoyConfigRevisionStatus) {
	*out = *in
	in.LastPublishedAt.DeepCopyInto(&out.LastPublishedAt)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigRevisionStatus.
func (in *EnvoyConfigRevisionStatus) DeepCopy() *EnvoyConfigRevisionStatus {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigRevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigSpec) DeepCopyInto(out *EnvoyConfigSpec) {
	*out = *in
	if in.Serialization != nil {
		in, out := &in.Serialization, &out.Serialization
		*out = new(string)
		**out = **in
	}
	if in.EnvoyAPI != nil {
		in, out := &in.EnvoyAPI, &out.EnvoyAPI
		*out = new(string)
		**out = **in
	}
	if in.EnvoyResources != nil {
		in, out := &in.EnvoyResources, &out.EnvoyResources
		*out = new(EnvoyResources)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigSpec.
func (in *EnvoyConfigSpec) DeepCopy() *EnvoyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyConfigStatus) DeepCopyInto(out *EnvoyConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigRevisions != nil {
		in, out := &in.ConfigRevisions, &out.ConfigRevisions
		*out = make([]ConfigRevisionRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyConfigStatus.
func (in *EnvoyConfigStatus) DeepCopy() *EnvoyConfigStatus {
	if in == nil {
		return nil
	}
	out := new(EnvoyConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyResource) DeepCopyInto(out *EnvoyResource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyResource.
func (in *EnvoyResource) DeepCopy() *EnvoyResource {
	if in == nil {
		return nil
	}
	out := new(EnvoyResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyResources) DeepCopyInto(out *EnvoyResources) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]EnvoyResource, len(*in))
		copy(*out, *in)
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]EnvoyResource, len(*in))
		copy(*out, *in)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]EnvoyResource, len(*in))
		copy(*out, *in)
	}
	if in.Listeners != nil {
		in, out := &in.Listeners, &out.Listeners
		*out = make([]EnvoyResource, len(*in))
		copy(*out, *in)
	}
	if in.Runtimes != nil {
		in, out := &in.Runtimes, &out.Runtimes
		*out = make([]EnvoyResource, len(*in))
		copy(*out, *in)
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]EnvoySecretResource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyResources.
func (in *EnvoyResources) DeepCopy() *EnvoyResources {
	if in == nil {
		return nil
	}
	out := new(EnvoyResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoySecretResource) DeepCopyInto(out *EnvoySecretResource) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoySecretResource.
func (in *EnvoySecretResource) DeepCopy() *EnvoySecretResource {
	if in == nil {
		return nil
	}
	out := new(EnvoySecretResource)
	in.DeepCopyInto(out)
	return out
}
