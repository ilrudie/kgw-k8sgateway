// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for test.multicluster.solo.io/v1alpha1 resources

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Generated Deepcopy methods for Test

func (in *Test) DeepCopyInto(out *Test) {
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)

	// deepcopy spec
	in.Spec.DeepCopyInto(&out.Spec)

	return
}

func (in *Test) DeepCopy() *Test {
	if in == nil {
		return nil
	}
	out := new(Test)
	in.DeepCopyInto(out)
	return out
}

func (in *Test) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *TestList) DeepCopyInto(out *TestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Test, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

func (in *TestList) DeepCopy() *TestList {
	if in == nil {
		return nil
	}
	out := new(TestList)
	in.DeepCopyInto(out)
	return out
}

func (in *TestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}