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

package config

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kind/pkg/kustomize"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Config contains cluster creation config
// This is the current internal config type used by cluster
// Other API versions can be converted to this struct with Convert()
type Config struct {
	metav1.TypeMeta

	// Image is the node image to use when running the cluster
	// TODO(bentheelder): split this into image and tag?
	Image string
	// KubeadmConfigPatches are applied to the generated kubeadm config as
	// strategic merge patches to `kustomize build` internally
	// https://github.com/kubernetes/community/blob/master/contributors/devel/strategic-merge-patch.md
	// This should be an inline yaml blob-string
	KubeadmConfigPatches []string
	// KubeadmConfigPatchesJSON6902 are applied to the generated kubeadm config
	// as patchesJson6902 to `kustomize build`
	KubeadmConfigPatchesJSON6902 []kustomize.PatchJSON6902
	// ControlPlane holds config for the control plane node
	ControlPlane *ControlPlane
}

// ControlPlane holds configurations specific to the control plane nodes
// (currently the only node).
type ControlPlane struct {
	// NodeLifecycle contains LifecycleHooks for phases of node provisioning
	NodeLifecycle *NodeLifecycle
}

// NodeLifecycle contains LifecycleHooks for phases of node provisioning
// Within each phase these hooks run in the order specified
type NodeLifecycle struct {
	// PreBoot hooks run before starting systemd
	PreBoot []LifecycleHook
	// PreKubeadm hooks run immediately before `kubeadm`
	PreKubeadm []LifecycleHook
	// PostKubeadm hooks run immediately after `kubeadm`
	PostKubeadm []LifecycleHook
	// PostSetup hooks run after any standard `kind` setup on the node
	PostSetup []LifecycleHook
}

// LifecycleHook represents a command to run at points in the node lifecycle
type LifecycleHook struct {
	// Name is used to improve logging (optional)
	Name string
	// Command is the command to run on the node
	Command []string
	// MustSucceed - if true then the hook / command failing will cause
	// cluster creation to fail, otherwise the error will just be logged and
	// the boot process will continue
	MustSucceed bool
}
