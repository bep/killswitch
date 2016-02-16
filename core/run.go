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

package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Run runs the exectutable defined in ctx and applies the conditionals.
func Run(ctx *Context, conditionals ...Conditional) error {

	// do a prestart check
	b, err := checkConditionals(ctx, conditionals...)

	if err != nil {
		return err
	}

	if !b {
		program := filepath.Base(ctx.Executable)
		err = notify(fmt.Sprintf("Cannot start %s", program), "Blocked by kill switch!")

		if err != nil {
			log.Println("Notify errored:", err)
		}
		return fmt.Errorf("Conditions evaluated to false before startup of %s", program)
	}

	var (
		cmd             *exec.Cmd
		condStopSignal  = make(chan bool, 1)
		execEndedSignal = make(chan bool, 1)
		ticker          *time.Ticker
	)

	cmd = exec.Command(ctx.Executable, strings.Split(ctx.Args, " ")...)

	go func(c *exec.Cmd, stop chan<- bool) {
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		err := c.Run()

		if err != nil {
			log.Printf("Command.Run: %s", err)
		}

		stop <- true
	}(cmd, execEndedSignal)

	interval := 5
	if ctx.Interval > 0 {
		interval = ctx.Interval
	}
	ticker = time.NewTicker(time.Duration(interval) * time.Second)

	go func(ticker *time.Ticker, stop chan<- bool) {
		for {
			select {
			case <-ticker.C:
				b, err := checkConditionals(ctx, conditionals...)
				if err != nil {
					log.Println("error: Conditionals check returned error:", err)
				}
				if !b {
					stop <- true
					return
				}
			}

		}
	}(ticker, condStopSignal)

	for {
		select {
		case <-condStopSignal:
			ticker.Stop()
			return killProcess(cmd)
		case <-execEndedSignal:
			log.Println("Program ended on its own ...")
			ticker.Stop()
			return nil
		}
	}

}

func killProcess(cmd *exec.Cmd) error {
	log.Println("Stopping ...")

	err := cmd.Process.Signal(os.Kill)
	if err != nil {
		// TODO(bep) try again? Other kill signal?
		return err
	}

	program := filepath.Base(cmd.Path)
	return notify(fmt.Sprintf("Killed %s", program), "Blocked by kill switch!")

}

func checkConditionals(ctx *Context, conditionals ...Conditional) (bool, error) {
	for _, c := range conditionals {
		b, err := c.Valid(ctx)

		if err != nil {
			return false, err
		}

		if !b {
			return false, nil
		}
	}

	return true, nil
}
