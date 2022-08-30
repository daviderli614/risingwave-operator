/*
 * Copyright 2022 Singularity Data
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package install

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"

	"github.com/singularity-data/risingwave-operator/pkg/command/context"
	"github.com/singularity-data/risingwave-operator/pkg/command/util"
)

const (
	UninstallExample = `  # uninstall the risingwave operator from the cluster
  kubectl rw uninstall
`
)

// UninstallOptions contains the options for the uninstallation command.
type UninstallOptions struct {
	genericclioptions.IOStreams
}

// NewUninstallOptions returns a UninstallOptions.
func NewUninstallOptions(streams genericclioptions.IOStreams) *UninstallOptions {
	return &UninstallOptions{
		IOStreams: streams,
	}
}

// NewUninstallCommand creates the uninstallation command of the operator.
func NewUninstallCommand(ctx *context.RWContext, streams genericclioptions.IOStreams) *cobra.Command {
	o := NewUninstallOptions(streams)

	cmd := &cobra.Command{
		Use:     "uninstall",
		Short:   "Uninstall the risingwave operator from the cluster",
		Long:    "Uninstall the risingwave operator from the cluster",
		Example: UninstallExample,
		Run: func(cmd *cobra.Command, args []string) {
			util.CheckErr(o.Complete(ctx, cmd, args))
			util.CheckErr(o.Run(ctx, cmd, args))
		},
	}

	return cmd
}

func (o *UninstallOptions) Complete(ctx *context.RWContext, cmd *cobra.Command, args []string) error {
	return nil
}

// Run will run the command as followed:
// 1. ensure the latest risingwave file exists.
// 2. delete the resource by the risingwave file.
// TODO: need record which versions installed.
// TODO: need to ensure no instances running before uninstall. Issue: https://github.com/singularity-data/risingwave-operator/issues/184
func (o *UninstallOptions) Run(ctx *context.RWContext, cmd *cobra.Command, args []string) error {
	yamlFile, err := Download(RisingWaveUrl, TemDir+"/operator")
	if err != nil {
		return err
	}

	err = ctx.Applier().UnApply(yamlFile)
	if err != nil {
		return err
	}
	return nil
}
