// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	kustomize "sigs.k8s.io/kind/pkg/kustomize"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerHandle) DeepCopyInto(out *ContainerHandle) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerHandle.
func (in *ContainerHandle) DeepCopy() *ContainerHandle {
	if in == nil {
		return nil
	}
	out := new(ContainerHandle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
	*out = *in
	if in.NodeLifecycle != nil {
		in, out := &in.NodeLifecycle, &out.NodeLifecycle
		*out = new(NodeLifecycle)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlane.
func (in *ControlPlane) DeepCopy() *ControlPlane {
	if in == nil {
		return nil
	}
	out := new(ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleHook) DeepCopyInto(out *LifecycleHook) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleHook.
func (in *LifecycleHook) DeepCopy() *LifecycleHook {
	if in == nil {
		return nil
	}
	out := new(LifecycleHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Node) DeepCopyInto(out *Node) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.KubeadmConfigPatches != nil {
		in, out := &in.KubeadmConfigPatches, &out.KubeadmConfigPatches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.KubeadmConfigPatchesJSON6902 != nil {
		in, out := &in.KubeadmConfigPatchesJSON6902, &out.KubeadmConfigPatchesJSON6902
		*out = make([]kustomize.PatchJSON6902, len(*in))
		copy(*out, *in)
	}
	if in.ControlPlane != nil {
		in, out := &in.ControlPlane, &out.ControlPlane
		*out = new(ControlPlane)
		(*in).DeepCopyInto(*out)
	}
	out.ContainerHandle = in.ContainerHandle
	return
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
func (in *NodeLifecycle) DeepCopyInto(out *NodeLifecycle) {
	*out = *in
	if in.PreBoot != nil {
		in, out := &in.PreBoot, &out.PreBoot
		*out = make([]LifecycleHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PreKubeadm != nil {
		in, out := &in.PreKubeadm, &out.PreKubeadm
		*out = make([]LifecycleHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostKubeadm != nil {
		in, out := &in.PostKubeadm, &out.PostKubeadm
		*out = make([]LifecycleHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostSetup != nil {
		in, out := &in.PostSetup, &out.PostSetup
		*out = make([]LifecycleHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLifecycle.
func (in *NodeLifecycle) DeepCopy() *NodeLifecycle {
	if in == nil {
		return nil
	}
	out := new(NodeLifecycle)
	in.DeepCopyInto(out)
	return out
}
