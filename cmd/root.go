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
	"github.com/spf13/cobra"
	"log"
)

var cfgFile string
var executable string
var execArgs string
var interval int
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "killswitch",
	Short: "A service to run on your PC to programs when certain conditions are met",
	Long:  `A service to run on your PC to programs when certain conditions are met.`,
}

func Execute() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		log.Println("Execute failed:", err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&executable, "exec", "e", "", "The program to execute")
	rootCmd.PersistentFlags().StringVarP(&execArgs, "args", "a", "", "The program argument list")
	rootCmd.PersistentFlags().IntVar(&interval, "interval", 5, "Interval between checks in seconds")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose logging")
}
