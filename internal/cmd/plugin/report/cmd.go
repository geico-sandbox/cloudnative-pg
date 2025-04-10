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

package report

import (
	"github.com/spf13/cobra"

	"github.com/cloudnative-pg/cloudnative-pg/internal/cmd/plugin"
)

// NewCmd creates the new "report" command
func NewCmd() *cobra.Command {
	reportCmd := &cobra.Command{
		Use:     "report operator/cluster",
		Short:   "Report on the operator or a cluster for troubleshooting",
		GroupID: plugin.GroupIDTroubleshooting,
	}

	reportCmd.AddCommand(operatorCmd())
	reportCmd.AddCommand(clusterCmd())

	return reportCmd
}
