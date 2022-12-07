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

type DevicesObservation struct {

	// The ID of the underlying entity this device references (i.e. the Linode's ID).
	EntityID *float64 `json:"entityId,omitempty" tf:"entity_id,omitempty"`

	// The ID of the Firewall.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// This Firewall's unique label.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// The type of Firewall Device.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The URL of the underlying entity this device references.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type DevicesParameters struct {
}

type FirewallObservation struct {

	// The devices associated with this firewall.
	Devices []DevicesObservation `json:"devices,omitempty" tf:"devices,omitempty"`

	// The ID of the Firewall.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The status of the Firewall.
	// The status of the firewall.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type FirewallParameters struct {

	// If true, the Firewall's rules are not enforced (defaults to false).
	// If true, the Firewall is inactive.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A firewall rule that specifies what inbound network traffic is allowed.
	// +kubebuilder:validation:Optional
	Inbound []InboundParameters `json:"inbound,omitempty" tf:"inbound,omitempty"`

	// The default behavior for inbound traffic. This setting can be overridden by updating the inbound.action property of the Firewall Rule. (ACCEPT, DROP)
	// The default behavior for inbound traffic. This setting can be overridden by updating the inbound.action property for an individual Firewall Rule.
	// +kubebuilder:validation:Required
	InboundPolicy *string `json:"inboundPolicy" tf:"inbound_policy,omitempty"`

	// This Firewall's unique label.
	// The label for the Firewall. For display purposes only. If no label is provided, a default will be assigned.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// A list of IDs of Linodes this Firewall should govern it's network traffic for.
	// The IDs of Linodes to apply this firewall to.
	// +crossplane:generate:reference:type=github.com/linode/provider-linode/apis/instance/v1alpha1.Instance
	// +kubebuilder:validation:Optional
	Linodes []*float64 `json:"linodes,omitempty" tf:"linodes,omitempty"`

	// References to Instance in instance to populate linodes.
	// +kubebuilder:validation:Optional
	LinodesRefs []v1.Reference `json:"linodesRefs,omitempty" tf:"-"`

	// Selector for a list of Instance in instance to populate linodes.
	// +kubebuilder:validation:Optional
	LinodesSelector *v1.Selector `json:"linodesSelector,omitempty" tf:"-"`

	// A firewall rule that specifies what outbound network traffic is allowed.
	// +kubebuilder:validation:Optional
	Outbound []OutboundParameters `json:"outbound,omitempty" tf:"outbound,omitempty"`

	// The default behavior for outbound traffic. This setting can be overridden by updating the outbound.action property for an individual Firewall Rule. (ACCEPT, DROP)
	// The default behavior for outbound traffic. This setting can be overridden by updating the outbound.action property for an individual Firewall Rule.
	// +kubebuilder:validation:Required
	OutboundPolicy *string `json:"outboundPolicy" tf:"outbound_policy,omitempty"`

	// A list of tags applied to the Kubernetes cluster. Tags are for organizational purposes only.
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type InboundObservation struct {
}

type InboundParameters struct {

	// Controls whether traffic is accepted or dropped by this rule (ACCEPT, DROP). Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.
	// Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// A list of IPv4 addresses or networks. Must be in IP/mask format.
	// A list of IP addresses, CIDR blocks, or 0.0.0.0/0 (to allow all) this rule applies to.
	// +kubebuilder:validation:Optional
	IPv4 []*string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// A list of IPv6 addresses or networks. Must be in IP/mask format.
	// A list of IPv6 addresses or networks this rule applies to.
	// +kubebuilder:validation:Optional
	IPv6 []*string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// This Firewall's unique label.
	// Used to identify this rule. For display purposes only.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
	// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
	// +kubebuilder:validation:Optional
	Ports *string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The network protocol this rule controls. (TCP, UDP, ICMP)
	// The network protocol this rule controls.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

type OutboundObservation struct {
}

type OutboundParameters struct {

	// Controls whether traffic is accepted or dropped by this rule (ACCEPT, DROP). Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.
	// Controls whether traffic is accepted or dropped by this rule. Overrides the Firewall’s inbound_policy if this is an inbound rule, or the outbound_policy if this is an outbound rule.
	// +kubebuilder:validation:Required
	Action *string `json:"action" tf:"action,omitempty"`

	// A list of IPv4 addresses or networks. Must be in IP/mask format.
	// A list of IP addresses, CIDR blocks, or 0.0.0.0/0 (to allow all) this rule applies to.
	// +kubebuilder:validation:Optional
	IPv4 []*string `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// A list of IPv6 addresses or networks. Must be in IP/mask format.
	// A list of IPv6 addresses or networks this rule applies to.
	// +kubebuilder:validation:Optional
	IPv6 []*string `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// This Firewall's unique label.
	// Used to identify this rule. For display purposes only.
	// +kubebuilder:validation:Required
	Label *string `json:"label" tf:"label,omitempty"`

	// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
	// A string representation of ports and/or port ranges (i.e. "443" or "80-90, 91").
	// +kubebuilder:validation:Optional
	Ports *string `json:"ports,omitempty" tf:"ports,omitempty"`

	// The network protocol this rule controls. (TCP, UDP, ICMP)
	// The network protocol this rule controls.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`
}

// FirewallSpec defines the desired state of Firewall
type FirewallSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallParameters `json:"forProvider"`
}

// FirewallStatus defines the observed state of Firewall.
type FirewallStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Firewall is the Schema for the Firewalls API. Manages a Linode Firewall.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Firewall struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec"`
	Status            FirewallStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallList contains a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Firewall `json:"items"`
}

// Repository type metadata.
var (
	Firewall_Kind             = "Firewall"
	Firewall_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Firewall_Kind}.String()
	Firewall_KindAPIVersion   = Firewall_Kind + "." + CRDGroupVersion.String()
	Firewall_GroupVersionKind = CRDGroupVersion.WithKind(Firewall_Kind)
)

func init() {
	SchemeBuilder.Register(&Firewall{}, &FirewallList{})
}
