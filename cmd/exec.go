// Copyright 2015 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/bep/killswitch/condition"
	"github.com/bep/killswitch/core"
	"github.com/spf13/cobra"
)

var heartbeatScript string

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Will kill your executable if your provided script exits with an error code",
	Long: `Will kill your executable if your provided script exits with an error code
	
The script (typically a shell script on *nix or a cmd- or bat-script on Windows) must exit with a non-0 exit-code 
to signal that the application under watch should be killed.

See /testfiles for example scripts for both *nix and Windows.

If the script is not present on the PATH, the full path must be provided in name.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(executable) == 0 {
			return userError{"Must provide an executable to watch. Note that this must either in PATH or set as a full path."}
		}
		fmt.Printf("Starting %s with kill switch script %s ...\n", executable, heartbeatScript)
		conditional := condition.NewExecConditional(heartbeatScript)
		ctx := &core.Context{Executable: executable, Args: execArgs, Interval: interval}
		return core.Run(ctx, conditional)
	},
}

func init() {
	rootCmd.AddCommand(execCmd)

	execCmd.Flags().StringVarP(&heartbeatScript, "name", "n", "", "The name of the script to use as heartbeat script. If not on PATH, the full path must be provided.")
}
