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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addon "sigs.k8s.io/controller-runtime/alpha/patterns/addon/pkg/apis/v1alpha1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IstioSpec defines the desired state of Istio
type IstioSpec struct {
	addon.CommonSpec
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// IstioStatus defines the observed state of Istio
type IstioStatus struct {
	addon.CommonStatus
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

var _ addon.CommonObject = &Istio{}

func (o *Istio) ComponentName() string {
	return "istio"
}

func (o *Istio) CommonSpec() addon.CommonSpec {
	return o.Spec.CommonSpec
}

func (o *Istio) GetCommonStatus() addon.CommonStatus {
	return o.Status.CommonStatus
}

func (o *Istio) SetCommonStatus(s addon.CommonStatus) {
	o.Status.CommonStatus = s
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Istio is the Schema for the istios API
// +k8s:openapi-gen=true
type Istio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IstioSpec   `json:"spec,omitempty"`
	Status IstioStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IstioList contains a list of Istio
type IstioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Istio `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Istio{}, &IstioList{})
}
