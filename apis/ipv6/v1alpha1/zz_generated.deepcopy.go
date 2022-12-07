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
func (in *IPv6Range) DeepCopyInto(out *IPv6Range) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6Range.
func (in *IPv6Range) DeepCopy() *IPv6Range {
	if in == nil {
		return nil
	}
	out := new(IPv6Range)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPv6Range) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6RangeList) DeepCopyInto(out *IPv6RangeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPv6Range, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6RangeList.
func (in *IPv6RangeList) DeepCopy() *IPv6RangeList {
	if in == nil {
		return nil
	}
	out := new(IPv6RangeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPv6RangeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6RangeObservation) DeepCopyInto(out *IPv6RangeObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.IsBGP != nil {
		in, out := &in.IsBGP, &out.IsBGP
		*out = new(bool)
		**out = **in
	}
	if in.Linodes != nil {
		in, out := &in.Linodes, &out.Linodes
		*out = make([]*float64, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(float64)
				**out = **in
			}
		}
	}
	if in.Range != nil {
		in, out := &in.Range, &out.Range
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6RangeObservation.
func (in *IPv6RangeObservation) DeepCopy() *IPv6RangeObservation {
	if in == nil {
		return nil
	}
	out := new(IPv6RangeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6RangeParameters) DeepCopyInto(out *IPv6RangeParameters) {
	*out = *in
	if in.LinodeID != nil {
		in, out := &in.LinodeID, &out.LinodeID
		*out = new(float64)
		**out = **in
	}
	if in.LinodeIDRef != nil {
		in, out := &in.LinodeIDRef, &out.LinodeIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LinodeIDSelector != nil {
		in, out := &in.LinodeIDSelector, &out.LinodeIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.PrefixLength != nil {
		in, out := &in.PrefixLength, &out.PrefixLength
		*out = new(float64)
		**out = **in
	}
	if in.RouteTarget != nil {
		in, out := &in.RouteTarget, &out.RouteTarget
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6RangeParameters.
func (in *IPv6RangeParameters) DeepCopy() *IPv6RangeParameters {
	if in == nil {
		return nil
	}
	out := new(IPv6RangeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6RangeSpec) DeepCopyInto(out *IPv6RangeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6RangeSpec.
func (in *IPv6RangeSpec) DeepCopy() *IPv6RangeSpec {
	if in == nil {
		return nil
	}
	out := new(IPv6RangeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6RangeStatus) DeepCopyInto(out *IPv6RangeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6RangeStatus.
func (in *IPv6RangeStatus) DeepCopy() *IPv6RangeStatus {
	if in == nil {
		return nil
	}
	out := new(IPv6RangeStatus)
	in.DeepCopyInto(out)
	return out
}
