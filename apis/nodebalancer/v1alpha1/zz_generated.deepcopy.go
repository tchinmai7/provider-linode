//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigList) DeepCopyInto(out *ConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Config, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigList.
func (in *ConfigList) DeepCopy() *ConfigList {
	if in == nil {
		return nil
	}
	out := new(ConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigObservation) DeepCopyInto(out *ConfigObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.NodeStatus != nil {
		in, out := &in.NodeStatus, &out.NodeStatus
		*out = make([]NodeStatusObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SSLCommonname != nil {
		in, out := &in.SSLCommonname, &out.SSLCommonname
		*out = new(string)
		**out = **in
	}
	if in.SSLFingerprint != nil {
		in, out := &in.SSLFingerprint, &out.SSLFingerprint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigObservation.
func (in *ConfigObservation) DeepCopy() *ConfigObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigParameters) DeepCopyInto(out *ConfigParameters) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = new(string)
		**out = **in
	}
	if in.CheckAttempts != nil {
		in, out := &in.CheckAttempts, &out.CheckAttempts
		*out = new(float64)
		**out = **in
	}
	if in.CheckBody != nil {
		in, out := &in.CheckBody, &out.CheckBody
		*out = new(string)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(float64)
		**out = **in
	}
	if in.CheckPassive != nil {
		in, out := &in.CheckPassive, &out.CheckPassive
		*out = new(bool)
		**out = **in
	}
	if in.CheckPath != nil {
		in, out := &in.CheckPath, &out.CheckPath
		*out = new(string)
		**out = **in
	}
	if in.CheckTimeout != nil {
		in, out := &in.CheckTimeout, &out.CheckTimeout
		*out = new(float64)
		**out = **in
	}
	if in.CipherSuite != nil {
		in, out := &in.CipherSuite, &out.CipherSuite
		*out = new(string)
		**out = **in
	}
	if in.NodebalancerID != nil {
		in, out := &in.NodebalancerID, &out.NodebalancerID
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.ProxyProtocol != nil {
		in, out := &in.ProxyProtocol, &out.ProxyProtocol
		*out = new(string)
		**out = **in
	}
	if in.SSLCertSecretRef != nil {
		in, out := &in.SSLCertSecretRef, &out.SSLCertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.SSLKeySecretRef != nil {
		in, out := &in.SSLKeySecretRef, &out.SSLKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.Stickiness != nil {
		in, out := &in.Stickiness, &out.Stickiness
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigParameters.
func (in *ConfigParameters) DeepCopy() *ConfigParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigSpec) DeepCopyInto(out *ConfigSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigSpec.
func (in *ConfigSpec) DeepCopy() *ConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigStatus) DeepCopyInto(out *ConfigStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigStatus.
func (in *ConfigStatus) DeepCopy() *ConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageList) DeepCopyInto(out *ImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Image, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageList.
func (in *ImageList) DeepCopy() *ImageList {
	if in == nil {
		return nil
	}
	out := new(ImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageObservation) DeepCopyInto(out *ImageObservation) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.Deprecated != nil {
		in, out := &in.Deprecated, &out.Deprecated
		*out = new(bool)
		**out = **in
	}
	if in.Expiry != nil {
		in, out := &in.Expiry, &out.Expiry
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsPublic != nil {
		in, out := &in.IsPublic, &out.IsPublic
		*out = new(bool)
		**out = **in
	}
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(float64)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Vendor != nil {
		in, out := &in.Vendor, &out.Vendor
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageObservation.
func (in *ImageObservation) DeepCopy() *ImageObservation {
	if in == nil {
		return nil
	}
	out := new(ImageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageParameters) DeepCopyInto(out *ImageParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DiskID != nil {
		in, out := &in.DiskID, &out.DiskID
		*out = new(float64)
		**out = **in
	}
	if in.FileHash != nil {
		in, out := &in.FileHash, &out.FileHash
		*out = new(string)
		**out = **in
	}
	if in.FilePath != nil {
		in, out := &in.FilePath, &out.FilePath
		*out = new(string)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(float64)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageParameters.
func (in *ImageParameters) DeepCopy() *ImageParameters {
	if in == nil {
		return nil
	}
	out := new(ImageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Node.
func (in *Node) DeepCopy() *Node {
	if in == nil {
		return nil
	}
	out := new(Node)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Node) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeList) DeepCopyInto(out *NodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Node, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeList.
func (in *NodeList) DeepCopy() *NodeList {
	if in == nil {
		return nil
	}
	out := new(NodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeObservation) DeepCopyInto(out *NodeObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeObservation.
func (in *NodeObservation) DeepCopy() *NodeObservation {
	if in == nil {
		return nil
	}
	out := new(NodeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeParameters) DeepCopyInto(out *NodeParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.ConfigID != nil {
		in, out := &in.ConfigID, &out.ConfigID
		*out = new(float64)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	if in.NodebalancerID != nil {
		in, out := &in.NodebalancerID, &out.NodebalancerID
		*out = new(float64)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeParameters.
func (in *NodeParameters) DeepCopy() *NodeParameters {
	if in == nil {
		return nil
	}
	out := new(NodeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSpec) DeepCopyInto(out *NodeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSpec.
func (in *NodeSpec) DeepCopy() *NodeSpec {
	if in == nil {
		return nil
	}
	out := new(NodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatus) DeepCopyInto(out *NodeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatus.
func (in *NodeStatus) DeepCopy() *NodeStatus {
	if in == nil {
		return nil
	}
	out := new(NodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatusObservation) DeepCopyInto(out *NodeStatusObservation) {
	*out = *in
	if in.Down != nil {
		in, out := &in.Down, &out.Down
		*out = new(float64)
		**out = **in
	}
	if in.Up != nil {
		in, out := &in.Up, &out.Up
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatusObservation.
func (in *NodeStatusObservation) DeepCopy() *NodeStatusObservation {
	if in == nil {
		return nil
	}
	out := new(NodeStatusObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeStatusParameters) DeepCopyInto(out *NodeStatusParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeStatusParameters.
func (in *NodeStatusParameters) DeepCopy() *NodeStatusParameters {
	if in == nil {
		return nil
	}
	out := new(NodeStatusParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Nodebalancer) DeepCopyInto(out *Nodebalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Nodebalancer.
func (in *Nodebalancer) DeepCopy() *Nodebalancer {
	if in == nil {
		return nil
	}
	out := new(Nodebalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Nodebalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerList) DeepCopyInto(out *NodebalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Nodebalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerList.
func (in *NodebalancerList) DeepCopy() *NodebalancerList {
	if in == nil {
		return nil
	}
	out := new(NodebalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodebalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerObservation) DeepCopyInto(out *NodebalancerObservation) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IPv4 != nil {
		in, out := &in.IPv4, &out.IPv4
		*out = new(string)
		**out = **in
	}
	if in.IPv6 != nil {
		in, out := &in.IPv6, &out.IPv6
		*out = new(string)
		**out = **in
	}
	if in.Transfer != nil {
		in, out := &in.Transfer, &out.Transfer
		*out = make([]TransferObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Updated != nil {
		in, out := &in.Updated, &out.Updated
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerObservation.
func (in *NodebalancerObservation) DeepCopy() *NodebalancerObservation {
	if in == nil {
		return nil
	}
	out := new(NodebalancerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerParameters) DeepCopyInto(out *NodebalancerParameters) {
	*out = *in
	if in.ClientConnThrottle != nil {
		in, out := &in.ClientConnThrottle, &out.ClientConnThrottle
		*out = new(float64)
		**out = **in
	}
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerParameters.
func (in *NodebalancerParameters) DeepCopy() *NodebalancerParameters {
	if in == nil {
		return nil
	}
	out := new(NodebalancerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerSpec) DeepCopyInto(out *NodebalancerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerSpec.
func (in *NodebalancerSpec) DeepCopy() *NodebalancerSpec {
	if in == nil {
		return nil
	}
	out := new(NodebalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodebalancerStatus) DeepCopyInto(out *NodebalancerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodebalancerStatus.
func (in *NodebalancerStatus) DeepCopy() *NodebalancerStatus {
	if in == nil {
		return nil
	}
	out := new(NodebalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferObservation) DeepCopyInto(out *TransferObservation) {
	*out = *in
	if in.In != nil {
		in, out := &in.In, &out.In
		*out = new(float64)
		**out = **in
	}
	if in.Out != nil {
		in, out := &in.Out, &out.Out
		*out = new(float64)
		**out = **in
	}
	if in.Total != nil {
		in, out := &in.Total, &out.Total
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferObservation.
func (in *TransferObservation) DeepCopy() *TransferObservation {
	if in == nil {
		return nil
	}
	out := new(TransferObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransferParameters) DeepCopyInto(out *TransferParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransferParameters.
func (in *TransferParameters) DeepCopy() *TransferParameters {
	if in == nil {
		return nil
	}
	out := new(TransferParameters)
	in.DeepCopyInto(out)
	return out
}
