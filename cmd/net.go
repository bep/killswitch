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
	"net"
)

var networkInterface string

var netinterfaceCmd = &cobra.Command{
	Use:   "net",
	Short: "Will kill your executable if a given network interface vanishes",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(executable) == 0 {
			return userError{"Must provide an executable to watch. Note that this must either in PATH or set as a full path."}
		}
		fmt.Printf("Starting %s with kill switch on interface %s ...\n", executable, networkInterface)
		conditional := condition.NewNetworkInterfaceConditional(networkInterface)
		ctx := &core.Context{Executable: executable, Args: execArgs, Interval: interval}
		return core.Run(ctx, conditional)
	},
}

var listInterfacesCmd = &cobra.Command{
	Use:   "list",
	Short: "List your network interfaces",

	RunE: func(cmd *cobra.Command, args []string) error {
		networkInterfaces, err := net.Interfaces()
		if err != nil {
			return err
		}
		fmt.Printf("Network interfaces (%d):\n", len(networkInterfaces))
		for i, in := range networkInterfaces {
			addrs, _ := in.Addrs()
			fmt.Println(i+1, in.Name, addrs)
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(netinterfaceCmd)
	netinterfaceCmd.AddCommand(listInterfacesCmd)

	netinterfaceCmd.Flags().StringVarP(&networkInterface, "name", "n", "", "The name of the network interface that must be present")
}
