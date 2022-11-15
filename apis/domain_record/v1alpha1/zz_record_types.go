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

type RecordObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RecordParameters struct {

	// The ID of the Domain to access.  Changing .
	// The ID of the Domain to access.
	// +kubebuilder:validation:Required
	DomainID *float64 `json:"domainId" tf:"domain_id,omitempty"`

	// The name of this Record. Setting this is invalid for SRV records as it is generated by the API. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	// The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address. Generated for SRV records.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The port this Record points to.
	// The port this Record points to.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// The priority of the target host. Lower values are preferred.
	// The priority of the target host. Lower values are preferred.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The protocol this Record's service communicates with. Only valid for SRV records.
	// The protocol this Record's service communicates with. Only valid for SRV records.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address. See all supported record types here. Changing .
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address.
	// +kubebuilder:validation:Required
	RecordType *string `json:"recordType" tf:"record_type,omitempty"`

	// The service this Record identified. Only valid for SRV records.
	// The service this Record identified. Only valid for SRV records.
	// +kubebuilder:validation:Optional
	Service *string `json:"service,omitempty" tf:"service,omitempty"`

	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +kubebuilder:validation:Optional
	TTLSec *float64 `json:"ttlSec,omitempty" tf:"ttl_sec,omitempty"`

	// The tag portion of a CAA record. It is invalid to set this on other record types.
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	// +kubebuilder:validation:Required
	Target *string `json:"target" tf:"target,omitempty"`

	// The relative weight of this Record. Higher values are preferred.
	// The relative weight of this Record. Higher values are preferred.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

// RecordSpec defines the desired state of Record
type RecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RecordParameters `json:"forProvider"`
}

// RecordStatus defines the observed state of Record.
type RecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Record is the Schema for the Records API. Manages a Linode Domain Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Record struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RecordSpec   `json:"spec"`
	Status            RecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RecordList contains a list of Records
type RecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Record `json:"items"`
}

// Repository type metadata.
var (
	Record_Kind             = "Record"
	Record_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Record_Kind}.String()
	Record_KindAPIVersion   = Record_Kind + "." + CRDGroupVersion.String()
	Record_GroupVersionKind = CRDGroupVersion.WithKind(Record_Kind)
)

func init() {
	SchemeBuilder.Register(&Record{}, &RecordList{})
}