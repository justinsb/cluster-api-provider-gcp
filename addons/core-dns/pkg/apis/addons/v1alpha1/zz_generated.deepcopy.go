// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreDNS) DeepCopyInto(out *CoreDNS) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreDNS.
func (in *CoreDNS) DeepCopy() *CoreDNS {
	if in == nil {
		return nil
	}
	out := new(CoreDNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoreDNS) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreDNSList) DeepCopyInto(out *CoreDNSList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoreDNS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreDNSList.
func (in *CoreDNSList) DeepCopy() *CoreDNSList {
	if in == nil {
		return nil
	}
	out := new(CoreDNSList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoreDNSList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreDNSSpec) DeepCopyInto(out *CoreDNSSpec) {
	*out = *in
	out.CommonSpec = in.CommonSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreDNSSpec.
func (in *CoreDNSSpec) DeepCopy() *CoreDNSSpec {
	if in == nil {
		return nil
	}
	out := new(CoreDNSSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreDNSStatus) DeepCopyInto(out *CoreDNSStatus) {
	*out = *in
	out.CommonStatus = in.CommonStatus
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreDNSStatus.
func (in *CoreDNSStatus) DeepCopy() *CoreDNSStatus {
	if in == nil {
		return nil
	}
	out := new(CoreDNSStatus)
	in.DeepCopyInto(out)
	return out
}
