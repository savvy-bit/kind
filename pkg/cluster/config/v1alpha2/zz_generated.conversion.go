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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "sigs.k8s.io/kind/pkg/cluster/config"
	kustomize "sigs.k8s.io/kind/pkg/kustomize"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Config)(nil), (*config.Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Config_To_config_Config(a.(*Config), b.(*config.Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Config)(nil), (*Config)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Config_To_v1alpha2_Config(a.(*config.Config), b.(*Config), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlane)(nil), (*config.ControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_ControlPlane_To_config_ControlPlane(a.(*ControlPlane), b.(*config.ControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ControlPlane)(nil), (*ControlPlane)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ControlPlane_To_v1alpha2_ControlPlane(a.(*config.ControlPlane), b.(*ControlPlane), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LifecycleHook)(nil), (*config.LifecycleHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_LifecycleHook_To_config_LifecycleHook(a.(*LifecycleHook), b.(*config.LifecycleHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LifecycleHook)(nil), (*LifecycleHook)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LifecycleHook_To_v1alpha2_LifecycleHook(a.(*config.LifecycleHook), b.(*LifecycleHook), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Node)(nil), (*config.Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Node_To_config_Node(a.(*Node), b.(*config.Node), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Node)(nil), (*Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Node_To_v1alpha2_Node(a.(*config.Node), b.(*Node), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeLifecycle)(nil), (*config.NodeLifecycle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_NodeLifecycle_To_config_NodeLifecycle(a.(*NodeLifecycle), b.(*config.NodeLifecycle), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeLifecycle)(nil), (*NodeLifecycle)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeLifecycle_To_v1alpha2_NodeLifecycle(a.(*config.NodeLifecycle), b.(*NodeLifecycle), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_Config_To_config_Config(in *Config, out *config.Config, s conversion.Scope) error {
	out.Nodes = *(*[]config.Node)(unsafe.Pointer(&in.Nodes))
	return nil
}

// Convert_v1alpha2_Config_To_config_Config is an autogenerated conversion function.
func Convert_v1alpha2_Config_To_config_Config(in *Config, out *config.Config, s conversion.Scope) error {
	return autoConvert_v1alpha2_Config_To_config_Config(in, out, s)
}

func autoConvert_config_Config_To_v1alpha2_Config(in *config.Config, out *Config, s conversion.Scope) error {
	out.Nodes = *(*[]Node)(unsafe.Pointer(&in.Nodes))
	// INFO: in.DerivedConfigData opted out of conversion generation
	return nil
}

// Convert_config_Config_To_v1alpha2_Config is an autogenerated conversion function.
func Convert_config_Config_To_v1alpha2_Config(in *config.Config, out *Config, s conversion.Scope) error {
	return autoConvert_config_Config_To_v1alpha2_Config(in, out, s)
}

func autoConvert_v1alpha2_ControlPlane_To_config_ControlPlane(in *ControlPlane, out *config.ControlPlane, s conversion.Scope) error {
	out.NodeLifecycle = (*config.NodeLifecycle)(unsafe.Pointer(in.NodeLifecycle))
	return nil
}

// Convert_v1alpha2_ControlPlane_To_config_ControlPlane is an autogenerated conversion function.
func Convert_v1alpha2_ControlPlane_To_config_ControlPlane(in *ControlPlane, out *config.ControlPlane, s conversion.Scope) error {
	return autoConvert_v1alpha2_ControlPlane_To_config_ControlPlane(in, out, s)
}

func autoConvert_config_ControlPlane_To_v1alpha2_ControlPlane(in *config.ControlPlane, out *ControlPlane, s conversion.Scope) error {
	out.NodeLifecycle = (*NodeLifecycle)(unsafe.Pointer(in.NodeLifecycle))
	return nil
}

// Convert_config_ControlPlane_To_v1alpha2_ControlPlane is an autogenerated conversion function.
func Convert_config_ControlPlane_To_v1alpha2_ControlPlane(in *config.ControlPlane, out *ControlPlane, s conversion.Scope) error {
	return autoConvert_config_ControlPlane_To_v1alpha2_ControlPlane(in, out, s)
}

func autoConvert_v1alpha2_LifecycleHook_To_config_LifecycleHook(in *LifecycleHook, out *config.LifecycleHook, s conversion.Scope) error {
	out.Name = in.Name
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.MustSucceed = in.MustSucceed
	return nil
}

// Convert_v1alpha2_LifecycleHook_To_config_LifecycleHook is an autogenerated conversion function.
func Convert_v1alpha2_LifecycleHook_To_config_LifecycleHook(in *LifecycleHook, out *config.LifecycleHook, s conversion.Scope) error {
	return autoConvert_v1alpha2_LifecycleHook_To_config_LifecycleHook(in, out, s)
}

func autoConvert_config_LifecycleHook_To_v1alpha2_LifecycleHook(in *config.LifecycleHook, out *LifecycleHook, s conversion.Scope) error {
	out.Name = in.Name
	out.Command = *(*[]string)(unsafe.Pointer(&in.Command))
	out.MustSucceed = in.MustSucceed
	return nil
}

// Convert_config_LifecycleHook_To_v1alpha2_LifecycleHook is an autogenerated conversion function.
func Convert_config_LifecycleHook_To_v1alpha2_LifecycleHook(in *config.LifecycleHook, out *LifecycleHook, s conversion.Scope) error {
	return autoConvert_config_LifecycleHook_To_v1alpha2_LifecycleHook(in, out, s)
}

func autoConvert_v1alpha2_Node_To_config_Node(in *Node, out *config.Node, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Role = config.NodeRole(in.Role)
	out.Image = in.Image
	out.KubeadmConfigPatches = *(*[]string)(unsafe.Pointer(&in.KubeadmConfigPatches))
	out.KubeadmConfigPatchesJSON6902 = *(*[]kustomize.PatchJSON6902)(unsafe.Pointer(&in.KubeadmConfigPatchesJSON6902))
	out.ControlPlane = (*config.ControlPlane)(unsafe.Pointer(in.ControlPlane))
	return nil
}

// Convert_v1alpha2_Node_To_config_Node is an autogenerated conversion function.
func Convert_v1alpha2_Node_To_config_Node(in *Node, out *config.Node, s conversion.Scope) error {
	return autoConvert_v1alpha2_Node_To_config_Node(in, out, s)
}

func autoConvert_config_Node_To_v1alpha2_Node(in *config.Node, out *Node, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Role = NodeRole(in.Role)
	out.Image = in.Image
	out.KubeadmConfigPatches = *(*[]string)(unsafe.Pointer(&in.KubeadmConfigPatches))
	out.KubeadmConfigPatchesJSON6902 = *(*[]kustomize.PatchJSON6902)(unsafe.Pointer(&in.KubeadmConfigPatchesJSON6902))
	out.ControlPlane = (*ControlPlane)(unsafe.Pointer(in.ControlPlane))
	return nil
}

// Convert_config_Node_To_v1alpha2_Node is an autogenerated conversion function.
func Convert_config_Node_To_v1alpha2_Node(in *config.Node, out *Node, s conversion.Scope) error {
	return autoConvert_config_Node_To_v1alpha2_Node(in, out, s)
}

func autoConvert_v1alpha2_NodeLifecycle_To_config_NodeLifecycle(in *NodeLifecycle, out *config.NodeLifecycle, s conversion.Scope) error {
	out.PreBoot = *(*[]config.LifecycleHook)(unsafe.Pointer(&in.PreBoot))
	out.PreKubeadm = *(*[]config.LifecycleHook)(unsafe.Pointer(&in.PreKubeadm))
	out.PostKubeadm = *(*[]config.LifecycleHook)(unsafe.Pointer(&in.PostKubeadm))
	out.PostSetup = *(*[]config.LifecycleHook)(unsafe.Pointer(&in.PostSetup))
	return nil
}

// Convert_v1alpha2_NodeLifecycle_To_config_NodeLifecycle is an autogenerated conversion function.
func Convert_v1alpha2_NodeLifecycle_To_config_NodeLifecycle(in *NodeLifecycle, out *config.NodeLifecycle, s conversion.Scope) error {
	return autoConvert_v1alpha2_NodeLifecycle_To_config_NodeLifecycle(in, out, s)
}

func autoConvert_config_NodeLifecycle_To_v1alpha2_NodeLifecycle(in *config.NodeLifecycle, out *NodeLifecycle, s conversion.Scope) error {
	out.PreBoot = *(*[]LifecycleHook)(unsafe.Pointer(&in.PreBoot))
	out.PreKubeadm = *(*[]LifecycleHook)(unsafe.Pointer(&in.PreKubeadm))
	out.PostKubeadm = *(*[]LifecycleHook)(unsafe.Pointer(&in.PostKubeadm))
	out.PostSetup = *(*[]LifecycleHook)(unsafe.Pointer(&in.PostSetup))
	return nil
}

// Convert_config_NodeLifecycle_To_v1alpha2_NodeLifecycle is an autogenerated conversion function.
func Convert_config_NodeLifecycle_To_v1alpha2_NodeLifecycle(in *config.NodeLifecycle, out *NodeLifecycle, s conversion.Scope) error {
	return autoConvert_config_NodeLifecycle_To_v1alpha2_NodeLifecycle(in, out, s)
}
