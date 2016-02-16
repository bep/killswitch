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
	"fmt"
	"github.com/bep/killswitch/core"
	"log"
	"os/exec"
	"sync/atomic"
)

type execConditional struct {
	execPath string

	//  used in tests
	appendCounter bool
	opCounter     uint64
}

func (c *execConditional) execAndArgs() (string, []string) {
	atomic.AddUint64(&c.opCounter, 1)
	if !c.appendCounter {
		return c.execPath, []string{}
	}
	return c.execPath, []string{fmt.Sprintf("%d", c.getOpCounter())}
}

func (c *execConditional) getOpCounter() uint64 {
	return atomic.LoadUint64(&c.opCounter)
}

func (c *execConditional) Valid(ctx *core.Context) (bool, error) {

	executable, args := c.execAndArgs()
	cmd := exec.Command(executable, args...)

	err := cmd.Run()

	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return false, nil
		}
		log.Printf("error: Verify Run: %s", err)
		return false, err

	}

	return true, nil
}

// NewExecConditional creates a Conditional that is valid unless the
// executable defined by execPath returns non-zero exit code.
func NewExecConditional(execPath string) core.Conditional {
	return &execConditional{execPath: execPath}
}

func newExecConditionalWithCounter(execPath string) *execConditional {
	return &execConditional{execPath: execPath, appendCounter: true}
}
