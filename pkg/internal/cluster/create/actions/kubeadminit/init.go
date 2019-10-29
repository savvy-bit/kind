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

// Package kubeadminit implements the kubeadm init action
package kubeadminit

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"sigs.k8s.io/kind/pkg/cluster/nodes"
	"sigs.k8s.io/kind/pkg/errors"
	"sigs.k8s.io/kind/pkg/exec"
	"sigs.k8s.io/kind/pkg/globals"

	"sigs.k8s.io/kind/pkg/cluster/nodeutils"

	"sigs.k8s.io/kind/pkg/internal/cluster/create/actions"
)

// kubeadmInitAction implements action for executing the kubadm init
// and a set of default post init operations like e.g. install the
// CNI network plugin.
type action struct{}

// NewAction returns a new action for kubeadm init
func NewAction() actions.Action {
	return &action{}
}

// Execute runs the action
func (a *action) Execute(ctx *actions.ActionContext) error {
	ctx.Status.Start("Starting control-plane 🕹️")
	defer ctx.Status.End(false)

	allNodes, err := ctx.Nodes()
	if err != nil {
		return err
	}

	// get the target node for this task
	// TODO: eliminate the concept of bootstrapcontrolplane node entirely
	// outside this method
	node, err := nodeutils.BootstrapControlPlaneNode(allNodes)
	if err != nil {
		return err
	}

	// run kubeadm
	cmd := node.Command(
		// init because this is the control plane node
		"kubeadm", "init",
		// preflight errors are expected, in particular for swap being enabled
		// TODO(bentheelder): limit the set of acceptable errors
		"--ignore-preflight-errors=all",
		// specify our generated config file
		"--config=/kind/kubeadm.conf",
		"--skip-token-print",
		// increase verbosity for debugging
		"--v=6",
	)
	lines, err := exec.CombinedOutputLines(cmd)
	globals.GetLogger().V(3).Info(strings.Join(lines, "\n"))
	if err != nil {
		return errors.Wrap(err, "failed to init node with kubeadm")
	}

	// copy some files to the other control plane nodes
	otherControlPlanes, err := nodeutils.SecondaryControlPlaneNodes(allNodes)
	if err != nil {
		return err
	}
	for _, otherNode := range otherControlPlanes {
		for _, file := range []string{
			// copy over admin config so we can use any control plane to get it later
			"/etc/kubernetes/admin.conf",
			// copy over certs
			"/etc/kubernetes/pki/ca.crt", "/etc/kubernetes/pki/ca.key",
			"/etc/kubernetes/pki/front-proxy-ca.crt", "/etc/kubernetes/pki/front-proxy-ca.key",
			"/etc/kubernetes/pki/sa.pub", "/etc/kubernetes/pki/sa.key",
			// TODO: if we gain external etcd support these will be
			// handled differently
			"/etc/kubernetes/pki/etcd/ca.crt", "/etc/kubernetes/pki/etcd/ca.key",
		} {
			if err := nodeutils.CopyNodeToNode(node, otherNode, file); err != nil {
				return errors.Wrap(err, "failed to copy admin kubeconfig")
			}
		}
	}

	// copies the kubeconfig files locally in order to make the cluster
	// usable with kubectl.
	// the kubeconfig file created by kubeadm internally to the node
	// must be modified in order to use the random host port reserved
	// for the API server and exposed by the node
	endpoint, err := ctx.ClusterContext.GetAPIServerEndpoint()
	if err != nil {
		return errors.Wrap(err, "failed to get api server endpoint from node")
	}
	kubeConfigPath := ctx.ClusterContext.KubeConfigPath()
	if err := writeKubeConfig(node, kubeConfigPath, endpoint); err != nil {
		return errors.Wrap(err, "failed to get kubeconfig from node")
	}

	// if we are only provisioning one node, remove the master taint
	// https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/#master-isolation
	if len(allNodes) == 1 {
		if err := node.Command(
			"kubectl", "--kubeconfig=/etc/kubernetes/admin.conf",
			"taint", "nodes", "--all", "node-role.kubernetes.io/master-",
		).Run(); err != nil {
			return errors.Wrap(err, "failed to remove master taint")
		}
	}

	// mark success
	ctx.Status.End(true)
	return nil
}

// matches kubeconfig server entry like:
//    server: https://172.17.0.2:6443
// which we rewrite to:
//    server: https://$ADDRESS:$PORT
var serverAddressRE = regexp.MustCompile(`^(\s+server:) https://.*:\d+$`)

// writeKubeConfig writes a fixed KUBECONFIG to dest
// this should only be called on a control plane node
// While copyng to the host machine the control plane address
// is replaced with local host and the control plane port with
// a randomly generated port reserved during node creation.
func writeKubeConfig(n nodes.Node, dest string, endpoint string) error {
	cmd := n.Command("cat", "/etc/kubernetes/admin.conf")
	lines, err := exec.CombinedOutputLines(cmd)
	if err != nil {
		return errors.Wrap(err, "failed to get kubeconfig from node")
	}

	// fix the config file, swapping out the server for the forwarded localhost:port
	var buff bytes.Buffer
	for _, line := range lines {
		match := serverAddressRE.FindStringSubmatch(line)
		if len(match) > 1 {
			line = fmt.Sprintf("%s https://%s", match[1], endpoint)
		}
		buff.WriteString(line)
		buff.WriteString("\n")
	}

	// create the directory to contain the KUBECONFIG file.
	// 0755 is taken from client-go's config handling logic: https://github.com/kubernetes/client-go/blob/5d107d4ebc00ee0ea606ad7e39fd6ce4b0d9bf9e/tools/clientcmd/loader.go#L412
	err = os.MkdirAll(filepath.Dir(dest), 0755)
	if err != nil {
		return errors.Wrap(err, "failed to create kubeconfig output directory")
	}

	return ioutil.WriteFile(dest, buff.Bytes(), 0600)
}
