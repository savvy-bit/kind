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

// Package kubeconfig provides utilites kind uses internally to manage
// kind cluster kubeconfigs
package kubeconfig

import (
	"bytes"

	"sigs.k8s.io/kind/pkg/cluster/nodeutils"
	"sigs.k8s.io/kind/pkg/errors"
	"sigs.k8s.io/kind/pkg/internal/cluster/context"

	// this package has slightly more generic kubeconfig helpers
	// and minimal dependencies on the rest of kind
	"sigs.k8s.io/kind/pkg/internal/cluster/kubeconfig/internal/kubeconfig"
)

// Export exports the kubeconfig given the cluster context and a path to write it to
// This will always be an external kubeconfig
func Export(ctx *context.Context, explicitPath string) error {
	cfg, err := get(ctx, true)
	if err != nil {
		return err
	}
	return kubeconfig.WriteMerged(cfg, explicitPath)
}

// Remove removes clusterName from the kubeconfig paths detected based on
// either explicitPath being set or $KUBECONFIG or $HOME/.kube/config, following
// the rules set by kubectl
// clusterName must identify a kind cluster.
func Remove(clusterName, explicitPath string) error {
	return kubeconfig.RemoveKIND(clusterName, explicitPath)
}

// Get returns the kubeconfig for the cluster
// external controls if the internal IP address is used or the host endpoint
func Get(ctx *context.Context, external bool) (string, error) {
	cfg, err := get(ctx, external)
	if err != nil {
		return "", err
	}
	b, err := kubeconfig.Encode(cfg)
	if err != nil {
		return "", err
	}
	return string(b), err
}

// ContextForCluster returns the context name for a kind cluster based on
// it's name. This key is used for all list entries of kind clusters
func ContextForCluster(kindClusterName string) string {
	return kubeconfig.KINDClusterKey(kindClusterName)
}

func get(ctx *context.Context, external bool) (*kubeconfig.Config, error) {
	// find a control plane node to get the kubeadm config from
	n, err := ctx.ListNodes()
	if err != nil {
		return nil, err
	}
	var buff bytes.Buffer
	nodes, err := nodeutils.ControlPlaneNodes(n)
	if err != nil {
		return nil, err
	}
	if len(nodes) < 1 {
		return nil, errors.New("could not locate any control plane nodes")
	}
	node := nodes[0]

	// grab kubeconfig version from the node
	if err := node.Command("cat", "/etc/kubernetes/admin.conf").SetStdout(&buff).Run(); err != nil {
		return nil, errors.Wrap(err, "failed to get cluster internal kubeconfig")
	}

	// if we're doing external we need to override the server endpoint
	server := ""
	if external {
		endpoint, err := ctx.GetAPIServerEndpoint()
		if err != nil {
			return nil, err
		}
		server = "https://" + endpoint
	}

	// actually encode
	return kubeconfig.KINDFromRawKubeadm(buff.String(), ctx.Name(), server)
}
