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

var networkInterface string

var netinterfaceCmd = &cobra.Command{
	Use:   "netinterface",
	Short: "Will kill your executable if a given network interface vanishes",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Starting %s with kill switch on interface %s ...\n", executable, networkInterface)
		conditional := condition.NewNetworkInterfaceConditional(networkInterface)
		ctx := &core.Context{Verbose: verbose, Executable: executable, Args: execArgs, Interval: interval}
		return core.Run(ctx, conditional)
	},
}

func init() {
	rootCmd.AddCommand(netinterfaceCmd)

	netinterfaceCmd.Flags().StringVarP(&networkInterface, "interface", "i", "", "The network interface that must be present")
}
