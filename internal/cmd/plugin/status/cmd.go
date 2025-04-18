/*
Copyright © contributors to CloudNativePG, established as
CloudNativePG a Series of LF Projects, LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package status

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cloudnative-pg/cloudnative-pg/internal/cmd/plugin"
)

// NewCmd create the new "status" subcommand
func NewCmd() *cobra.Command {
	statusCmd := &cobra.Command{
		Use:     "status CLUSTER",
		Short:   "Get the status of a PostgreSQL cluster",
		Args:    plugin.RequiresArguments(1),
		GroupID: plugin.GroupIDDatabase,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			if strings.HasPrefix(toComplete, "-") {
				fmt.Printf("%+v\n", toComplete)
			}
			return plugin.CompleteClusters(cmd.Context(), args, toComplete), cobra.ShellCompDirectiveNoFileComp
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			clusterName := args[0]

			verbose, _ := cmd.Flags().GetCount("verbose")
			output, _ := cmd.Flags().GetString("output")

			return Status(ctx, clusterName, verbose, plugin.OutputFormat(output))
		},
	}

	statusCmd.Flags().CountP(
		"verbose", "v", "Increase verbosity to display more information")
	statusCmd.Flags().StringP(
		"output", "o", "text", "Output format. One of text|json")

	return statusCmd
}
