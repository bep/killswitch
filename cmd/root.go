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
	"os"
)

type userError struct {
	s string
}

func (u userError) Error() string {
	return u.s
}

var cfgFile string
var executable string
var execArgs string
var interval int
var verbose bool

var rootCmd = &cobra.Command{
	Use:   "killswitch",
	Short: "Wrap your sensitive application with a kill switch",
	Long:  `Wrap your sensitive application with a kill switch.`,
}

func Execute() {

	rootCmd.SetOutput(logWriter)
	rootCmd.SilenceUsage = true

	if c, err := rootCmd.ExecuteC(); err != nil {
		if _, ok := err.(userError); ok {
			c.Println("")
			c.Println(c.UsageString())
		}

		os.Exit(-1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&executable, "exec", "e", "", "The program to watch")
	rootCmd.PersistentFlags().StringVarP(&execArgs, "args", "a", "", "The program argument list")
	rootCmd.PersistentFlags().IntVar(&interval, "interval", 5, "Interval between checks in seconds")
}
