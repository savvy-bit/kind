// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

package v1alpha3

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "sigs.k8s.io/kind/pkg/cluster/config"
	cri "sigs.k8s.io/kind/pkg/container/cri"
	kustomize "sigs.k8s.io/kind/pkg/kustomize"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Cluster)(nil), (*config.Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Cluster_To_config_Cluster(a.(*Cluster), b.(*config.Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Cluster)(nil), (*Cluster)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Cluster_To_v1alpha3_Cluster(a.(*config.Cluster), b.(*Cluster), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Networking)(nil), (*config.Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Networking_To_config_Networking(a.(*Networking), b.(*config.Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Networking)(nil), (*Networking)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Networking_To_v1alpha3_Networking(a.(*config.Networking), b.(*Networking), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Node)(nil), (*config.Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha3_Node_To_config_Node(a.(*Node), b.(*config.Node), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Node)(nil), (*Node)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Node_To_v1alpha3_Node(a.(*config.Node), b.(*Node), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha3_Cluster_To_config_Cluster(in *Cluster, out *config.Cluster, s conversion.Scope) error {
	out.Nodes = *(*[]config.Node)(unsafe.Pointer(&in.Nodes))
	if err := Convert_v1alpha3_Networking_To_config_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubeadmConfigPatches = *(*[]string)(unsafe.Pointer(&in.KubeadmConfigPatches))
	out.KubeadmConfigPatchesJSON6902 = *(*[]kustomize.PatchJSON6902)(unsafe.Pointer(&in.KubeadmConfigPatchesJSON6902))
	return nil
}

// Convert_v1alpha3_Cluster_To_config_Cluster is an autogenerated conversion function.
func Convert_v1alpha3_Cluster_To_config_Cluster(in *Cluster, out *config.Cluster, s conversion.Scope) error {
	return autoConvert_v1alpha3_Cluster_To_config_Cluster(in, out, s)
}

func autoConvert_config_Cluster_To_v1alpha3_Cluster(in *config.Cluster, out *Cluster, s conversion.Scope) error {
	out.Nodes = *(*[]Node)(unsafe.Pointer(&in.Nodes))
	if err := Convert_config_Networking_To_v1alpha3_Networking(&in.Networking, &out.Networking, s); err != nil {
		return err
	}
	out.KubeadmConfigPatches = *(*[]string)(unsafe.Pointer(&in.KubeadmConfigPatches))
	out.KubeadmConfigPatchesJSON6902 = *(*[]kustomize.PatchJSON6902)(unsafe.Pointer(&in.KubeadmConfigPatchesJSON6902))
	return nil
}

// Convert_config_Cluster_To_v1alpha3_Cluster is an autogenerated conversion function.
func Convert_config_Cluster_To_v1alpha3_Cluster(in *config.Cluster, out *Cluster, s conversion.Scope) error {
	return autoConvert_config_Cluster_To_v1alpha3_Cluster(in, out, s)
}

func autoConvert_v1alpha3_Networking_To_config_Networking(in *Networking, out *config.Networking, s conversion.Scope) error {
	out.APIServerPort = in.APIServerPort
	out.APIServerAddress = in.APIServerAddress
	return nil
}

// Convert_v1alpha3_Networking_To_config_Networking is an autogenerated conversion function.
func Convert_v1alpha3_Networking_To_config_Networking(in *Networking, out *config.Networking, s conversion.Scope) error {
	return autoConvert_v1alpha3_Networking_To_config_Networking(in, out, s)
}

func autoConvert_config_Networking_To_v1alpha3_Networking(in *config.Networking, out *Networking, s conversion.Scope) error {
	out.APIServerPort = in.APIServerPort
	out.APIServerAddress = in.APIServerAddress
	return nil
}

// Convert_config_Networking_To_v1alpha3_Networking is an autogenerated conversion function.
func Convert_config_Networking_To_v1alpha3_Networking(in *config.Networking, out *Networking, s conversion.Scope) error {
	return autoConvert_config_Networking_To_v1alpha3_Networking(in, out, s)
}

func autoConvert_v1alpha3_Node_To_config_Node(in *Node, out *config.Node, s conversion.Scope) error {
	out.Role = config.NodeRole(in.Role)
	out.Image = in.Image
	out.ExtraMounts = *(*[]cri.Mount)(unsafe.Pointer(&in.ExtraMounts))
	return nil
}

// Convert_v1alpha3_Node_To_config_Node is an autogenerated conversion function.
func Convert_v1alpha3_Node_To_config_Node(in *Node, out *config.Node, s conversion.Scope) error {
	return autoConvert_v1alpha3_Node_To_config_Node(in, out, s)
}

func autoConvert_config_Node_To_v1alpha3_Node(in *config.Node, out *Node, s conversion.Scope) error {
	out.Role = NodeRole(in.Role)
	out.Image = in.Image
	out.ExtraMounts = *(*[]cri.Mount)(unsafe.Pointer(&in.ExtraMounts))
	return nil
}

// Convert_config_Node_To_v1alpha3_Node is an autogenerated conversion function.
func Convert_config_Node_To_v1alpha3_Node(in *config.Node, out *Node, s conversion.Scope) error {
	return autoConvert_config_Node_To_v1alpha3_Node(in, out, s)
}
