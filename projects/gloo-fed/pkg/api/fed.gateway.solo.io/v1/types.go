// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1

import (
	. "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.gateway.solo.io/v1/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FederatedGateway
var FederatedGatewayGVK = schema.GroupVersionKind{
	Group:   "fed.gateway.solo.io",
	Version: "v1",
	Kind:    "FederatedGateway",
}

// FederatedGateway is the Schema for the federatedGateway API
type FederatedGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedGatewaySpec   `json:"spec,omitempty"`
	Status FederatedGatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FederatedGateway) GVK() schema.GroupVersionKind {
	return FederatedGatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedGatewayList contains a list of FederatedGateway
type FederatedGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedGateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FederatedMatchableHttpGateway
var FederatedMatchableHttpGatewayGVK = schema.GroupVersionKind{
	Group:   "fed.gateway.solo.io",
	Version: "v1",
	Kind:    "FederatedMatchableHttpGateway",
}

// FederatedMatchableHttpGateway is the Schema for the federatedMatchableHttpGateway API
type FederatedMatchableHttpGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedMatchableHttpGatewaySpec   `json:"spec,omitempty"`
	Status FederatedMatchableHttpGatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FederatedMatchableHttpGateway) GVK() schema.GroupVersionKind {
	return FederatedMatchableHttpGatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedMatchableHttpGatewayList contains a list of FederatedMatchableHttpGateway
type FederatedMatchableHttpGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedMatchableHttpGateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FederatedVirtualService
var FederatedVirtualServiceGVK = schema.GroupVersionKind{
	Group:   "fed.gateway.solo.io",
	Version: "v1",
	Kind:    "FederatedVirtualService",
}

// FederatedVirtualService is the Schema for the federatedVirtualService API
type FederatedVirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedVirtualServiceSpec   `json:"spec,omitempty"`
	Status FederatedVirtualServiceStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FederatedVirtualService) GVK() schema.GroupVersionKind {
	return FederatedVirtualServiceGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedVirtualServiceList contains a list of FederatedVirtualService
type FederatedVirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedVirtualService `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for FederatedRouteTable
var FederatedRouteTableGVK = schema.GroupVersionKind{
	Group:   "fed.gateway.solo.io",
	Version: "v1",
	Kind:    "FederatedRouteTable",
}

// FederatedRouteTable is the Schema for the federatedRouteTable API
type FederatedRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FederatedRouteTableSpec   `json:"spec,omitempty"`
	Status FederatedRouteTableStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (FederatedRouteTable) GVK() schema.GroupVersionKind {
	return FederatedRouteTableGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FederatedRouteTableList contains a list of FederatedRouteTable
type FederatedRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FederatedRouteTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FederatedGateway{}, &FederatedGatewayList{})
	SchemeBuilder.Register(&FederatedMatchableHttpGateway{}, &FederatedMatchableHttpGatewayList{})
	SchemeBuilder.Register(&FederatedVirtualService{}, &FederatedVirtualServiceList{})
	SchemeBuilder.Register(&FederatedRouteTable{}, &FederatedRouteTableList{})
}
