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

type StackscriptInitParameters struct {

	// A description for the StackScript.
	// A description for the StackScript.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of Image IDs representing the Images that this StackScript is compatible for deploying with. any/all indicates that all available image distributions, including private images, are accepted. Currently private image IDs are not supported.
	// An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
	Images []*string `json:"images,omitempty" tf:"images,omitempty"`

	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private. Changing
	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// The StackScript's label is for display purposes only.
	// The StackScript's label is for display purposes only.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// This field allows you to add notes for the set of revisions made to this StackScript.
	// This field allows you to add notes for the set of revisions made to this StackScript.
	RevNote *string `json:"revNote,omitempty" tf:"rev_note,omitempty"`

	// The script to execute when provisioning a new Linode with this StackScript.
	// The script to execute when provisioning a new Linode with this StackScript.
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type StackscriptObservation struct {

	// The date this StackScript was created.
	// The date this StackScript was created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// Count of currently active, deployed Linodes created from this StackScript.
	// Count of currently active, deployed Linodes created from this StackScript.
	DeploymentsActive *float64 `json:"deploymentsActive,omitempty" tf:"deployments_active,omitempty"`

	// The total number of times this StackScript has been deployed.
	// The total number of times this StackScript has been deployed.
	DeploymentsTotal *float64 `json:"deploymentsTotal,omitempty" tf:"deployments_total,omitempty"`

	// A description for the StackScript.
	// A description for the StackScript.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of Image IDs representing the Images that this StackScript is compatible for deploying with. any/all indicates that all available image distributions, including private images, are accepted. Currently private image IDs are not supported.
	// An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
	Images []*string `json:"images,omitempty" tf:"images,omitempty"`

	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private. Changing
	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// The StackScript's label is for display purposes only.
	// The StackScript's label is for display purposes only.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// This field allows you to add notes for the set of revisions made to this StackScript.
	// This field allows you to add notes for the set of revisions made to this StackScript.
	RevNote *string `json:"revNote,omitempty" tf:"rev_note,omitempty"`

	// The script to execute when provisioning a new Linode with this StackScript.
	// The script to execute when provisioning a new Linode with this StackScript.
	Script *string `json:"script,omitempty" tf:"script,omitempty"`

	// The date this StackScript was updated.
	// The date this StackScript was updated.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`

	// This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.
	// This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment.
	UserDefinedFields []UserDefinedFieldsObservation `json:"userDefinedFields,omitempty" tf:"user_defined_fields,omitempty"`

	// The Gravatar ID for the User who created the StackScript.
	// The Gravatar ID for the User who created the StackScript.
	UserGravatarID *string `json:"userGravatarId,omitempty" tf:"user_gravatar_id,omitempty"`

	// The User who created the StackScript.
	// The User who created the StackScript.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type StackscriptParameters struct {

	// A description for the StackScript.
	// A description for the StackScript.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A set of Image IDs representing the Images that this StackScript is compatible for deploying with. any/all indicates that all available image distributions, including private images, are accepted. Currently private image IDs are not supported.
	// An array of Image IDs representing the Images that this StackScript is compatible for deploying with.
	// +kubebuilder:validation:Optional
	Images []*string `json:"images,omitempty" tf:"images,omitempty"`

	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private. Changing
	// This determines whether other users can use your StackScript. Once a StackScript is made public, it cannot be made private.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// The StackScript's label is for display purposes only.
	// The StackScript's label is for display purposes only.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// This field allows you to add notes for the set of revisions made to this StackScript.
	// This field allows you to add notes for the set of revisions made to this StackScript.
	// +kubebuilder:validation:Optional
	RevNote *string `json:"revNote,omitempty" tf:"rev_note,omitempty"`

	// The script to execute when provisioning a new Linode with this StackScript.
	// The script to execute when provisioning a new Linode with this StackScript.
	// +kubebuilder:validation:Optional
	Script *string `json:"script,omitempty" tf:"script,omitempty"`
}

type UserDefinedFieldsInitParameters struct {
}

type UserDefinedFieldsObservation struct {

	// The default value. If not specified, this value will be used.
	Default *string `json:"default,omitempty" tf:"default,omitempty"`

	// An example value for the field.
	Example *string `json:"example,omitempty" tf:"example,omitempty"`

	// The StackScript's label is for display purposes only.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	// A list of acceptable values for the field in any quantity, combination or order.
	ManyOf *string `json:"manyOf,omitempty" tf:"many_of,omitempty"`

	// The name of the field.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A list of acceptable single values for the field.
	OneOf *string `json:"oneOf,omitempty" tf:"one_of,omitempty"`
}

type UserDefinedFieldsParameters struct {
}

// StackscriptSpec defines the desired state of Stackscript
type StackscriptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackscriptParameters `json:"forProvider"`
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
	InitProvider StackscriptInitParameters `json:"initProvider,omitempty"`
}

// StackscriptStatus defines the observed state of Stackscript.
type StackscriptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackscriptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stackscript is the Schema for the Stackscripts API. Manages a Linode StackScript.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,linode}
type Stackscript struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.images) || (has(self.initProvider) && has(self.initProvider.images))",message="spec.forProvider.images is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.label) || (has(self.initProvider) && has(self.initProvider.label))",message="spec.forProvider.label is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.script) || (has(self.initProvider) && has(self.initProvider.script))",message="spec.forProvider.script is a required parameter"
	Spec   StackscriptSpec   `json:"spec"`
	Status StackscriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackscriptList contains a list of Stackscripts
type StackscriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stackscript `json:"items"`
}

// Repository type metadata.
var (
	Stackscript_Kind             = "Stackscript"
	Stackscript_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stackscript_Kind}.String()
	Stackscript_KindAPIVersion   = Stackscript_Kind + "." + CRDGroupVersion.String()
	Stackscript_GroupVersionKind = CRDGroupVersion.WithKind(Stackscript_Kind)
)

func init() {
	SchemeBuilder.Register(&Stackscript{}, &StackscriptList{})
}
