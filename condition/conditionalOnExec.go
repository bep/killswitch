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

package condition

import (
	"github.com/bep/killswitch/core"
	"log"
	"os"
	"os/exec"
)

type execConditional struct {
	execPath string
}

func (n execConditional) Valid(ctx *core.Context) (bool, error) {

	c := exec.Command(n.execPath)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return false, nil
		}
		log.Printf("error: Verify Run: %s", err)
		return false, nil
	}

	return true, nil
}

func NewExecConditional(execPath string) core.Conditional {
	return execConditional{execPath: execPath}
}
