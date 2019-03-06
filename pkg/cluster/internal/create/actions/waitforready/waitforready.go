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

// Package waitforready implements the wait for ready action
package waitforready

import (
	"fmt"
	"time"

	"sigs.k8s.io/kind/pkg/cluster/internal/create/actions"
	"sigs.k8s.io/kind/pkg/cluster/nodes"
)

// Action implements an action for waiting for the cluster to be ready
type Action struct {
	waitTime time.Duration
}

// NewAction returns a new action for waiting for the cluster to be ready
func NewAction(waitTime time.Duration) actions.Action {
	return &Action{
		waitTime: waitTime,
	}
}

// Execute runs the action
func (a *Action) Execute(ctx *actions.ActionContext) error {
	ctx.Status.Start(
		fmt.Sprintf(
			"Waiting %s for the cluster to be ready ⏳",
			a.waitTime.Round(time.Second).String(),
		),
	)

	// get the bootstrap control plane node to use to check cluster status
	allNodes, err := ctx.Nodes()
	if err != nil {
		return err
	}
	node, err := nodes.BootstrapControlPlaneNode(allNodes)
	if err != nil {
		return err
	}

	// Wait for the nodes to reach Ready status.
	isReady := nodes.WaitForReady(node, time.Now().Add(a.waitTime))
	if !isReady {
		ctx.Status.End(false)
		fmt.Println("WARNING: Timed out waiting for the nodes to be ready.")
		return nil
	}

	// mark success
	ctx.Status.End(true)
	return nil
}
