// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NodeInitParameters struct {

	// The private IP Address where this backend can be reached. This must be a private IP address.
	// The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The label of the Linode NodeBalancer Node. This is for display purposes only.
	// The label for this node. This is for display purposes only.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The mode this NodeBalancer should use when sending traffic to this backend. If set to accept this backend is accepting traffic. If set to reject this backend will not receive traffic. If set to drain this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (accept, reject, drain, backup)
	// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. If set to `backup` this backend will only accept traffic if all other nodes are down.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).
	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type NodeObservation struct {

	// The private IP Address where this backend can be reached. This must be a private IP address.
	// The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ID of the NodeBalancerConfig to access.
	// The ID of the NodeBalancerConfig to access.
	ConfigID *float64 `json:"configId,omitempty" tf:"config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The label of the Linode NodeBalancer Node. This is for display purposes only.
	// The label for this node. This is for display purposes only.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The mode this NodeBalancer should use when sending traffic to this backend. If set to accept this backend is accepting traffic. If set to reject this backend will not receive traffic. If set to drain this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (accept, reject, drain, backup)
	// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. If set to `backup` this backend will only accept traffic if all other nodes are down.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the NodeBalancer to access.
	// The ID of the NodeBalancer to access.
	NodebalancerID *float64 `json:"nodebalancerId,omitempty" tf:"nodebalancer_id,omitempty"`

	// The current status of this node, based on the configured checks of its NodeBalancer Config. (unknown, UP, DOWN).
	// The current status of this node, based on the configured checks of its NodeBalancer Config. (unknown, UP, DOWN)
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).
	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type NodeParameters struct {

	// The private IP Address where this backend can be reached. This must be a private IP address.
	// The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The ID of the NodeBalancerConfig to access.
	// The ID of the NodeBalancerConfig to access.
	// +crossplane:generate:reference:type=Config
	// +kubebuilder:validation:Optional
	ConfigID *float64 `json:"configId,omitempty" tf:"config_id,omitempty"`

	// Reference to a Config to populate configId.
	// +kubebuilder:validation:Optional
	ConfigIDRef *v1.Reference `json:"configIdRef,omitempty" tf:"-"`

	// Selector for a Config to populate configId.
	// +kubebuilder:validation:Optional
	ConfigIDSelector *v1.Selector `json:"configIdSelector,omitempty" tf:"-"`

	// The label of the Linode NodeBalancer Node. This is for display purposes only.
	// The label for this node. This is for display purposes only.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The mode this NodeBalancer should use when sending traffic to this backend. If set to accept this backend is accepting traffic. If set to reject this backend will not receive traffic. If set to drain this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. (accept, reject, drain, backup)
	// The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it. If set to `backup` this backend will only accept traffic if all other nodes are down.
	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// The ID of the NodeBalancer to access.
	// The ID of the NodeBalancer to access.
	// +crossplane:generate:reference:type=Nodebalancer
	// +kubebuilder:validation:Optional
	NodebalancerID *float64 `json:"nodebalancerId,omitempty" tf:"nodebalancer_id,omitempty"`

	// Reference to a Nodebalancer to populate nodebalancerId.
	// +kubebuilder:validation:Optional
	NodebalancerIDRef *v1.Reference `json:"nodebalancerIdRef,omitempty" tf:"-"`

	// Selector for a Nodebalancer to populate nodebalancerId.
	// +kubebuilder:validation:Optional
	NodebalancerIDSelector *v1.Selector `json:"nodebalancerIdSelector,omitempty" tf:"-"`

	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255).
	// Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// NodeSpec defines the desired state of Node
type NodeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NodeInitParameters `json:"initProvider,omitempty"`
}

// NodeStatus defines the observed state of Node.
type NodeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Node is the Schema for the Nodes API. Manages a Linode NodeBalancer Node.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Node struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.address) || (has(self.initProvider) && has(self.initProvider.address))",message="spec.forProvider.address is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	Spec   NodeSpec   `json:"spec"`
	Status NodeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeList contains a list of Nodes
type NodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Node `json:"items"`
}

// Repository type metadata.
var (
	Node_Kind             = "Node"
	Node_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Node_Kind}.String()
	Node_KindAPIVersion   = Node_Kind + "." + CRDGroupVersion.String()
	Node_GroupVersionKind = CRDGroupVersion.WithKind(Node_Kind)
)

func init() {
	SchemeBuilder.Register(&Node{}, &NodeList{})
}
