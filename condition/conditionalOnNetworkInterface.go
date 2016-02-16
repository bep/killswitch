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
	"net"
)

type networkInterfaceConditional struct {
	name string
}

func (n networkInterfaceConditional) Valid(ctx *core.Context) (bool, error) {
	networkInterfaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, in := range networkInterfaces {
		if in.Name == n.name {
			return true, nil
		}

	}
	return false, nil
}

// NewNetworkInterfaceConditional creates a Conditional that is valid
// as long as the named network interface is present.
func NewNetworkInterfaceConditional(name string) core.Conditional {
	return networkInterfaceConditional{name: name}
}
