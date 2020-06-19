// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirFactoryService) DeepCopyInto(out *AirFactoryService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirFactoryService.
func (in *AirFactoryService) DeepCopy() *AirFactoryService {
	if in == nil {
		return nil
	}
	out := new(AirFactoryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AirFactoryService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirFactoryServiceList) DeepCopyInto(out *AirFactoryServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AirFactoryService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirFactoryServiceList.
func (in *AirFactoryServiceList) DeepCopy() *AirFactoryServiceList {
	if in == nil {
		return nil
	}
	out := new(AirFactoryServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AirFactoryServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirFactoryServiceSpec) DeepCopyInto(out *AirFactoryServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirFactoryServiceSpec.
func (in *AirFactoryServiceSpec) DeepCopy() *AirFactoryServiceSpec {
	if in == nil {
		return nil
	}
	out := new(AirFactoryServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AirFactoryServiceStatus) DeepCopyInto(out *AirFactoryServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AirFactoryServiceStatus.
func (in *AirFactoryServiceStatus) DeepCopy() *AirFactoryServiceStatus {
	if in == nil {
		return nil
	}
	out := new(AirFactoryServiceStatus)
	in.DeepCopyInto(out)
	return out
}
