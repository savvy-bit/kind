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

// Package kubeconfigpath implements the `kubeconfig-path` command
package kubeconfigpath

import (
	"fmt"

	"github.com/spf13/cobra"

	"sigs.k8s.io/kind/pkg/cluster"
)

type flags struct {
	Name string
}

// NewCommand returns a new cobra.Command for getting the kubeconfig path
func NewCommand() *cobra.Command {
	flags := &flags{}
	cmd := &cobra.Command{
		// TODO(bentheelder): more detailed usage
		Use:   "kubeconfig-path",
		Short: "prints the default kubeconfig path for the kind cluster by --name",
		Long:  "prints the default kubeconfig path for the kind cluster by --name",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runE(flags, cmd, args)
		},
	}
	cmd.Flags().StringVar(
		&flags.Name,
		"name",
		"1",
		"the cluster context name",
	)
	return cmd
}

func runE(flags *flags, cmd *cobra.Command, args []string) error {
	ctx, err := cluster.NewContext(flags.Name)
	if err != nil {
		return fmt.Errorf("failed to create cluster context! %v", err)
	}
	fmt.Println(ctx.KubeConfigPath())
	return nil
}
