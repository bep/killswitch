// +build !windows

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

package main

import (
	"github.com/bep/killswitch/cmd"
	"io"
	"log"
	"log/syslog"
	"os"
)

func init() {
	syslogWriter, e := syslog.New(syslog.LOG_NOTICE, "killswitch")
	logWriter := io.MultiWriter(syslogWriter, os.Stdout)

	if e == nil {
		log.SetOutput(logWriter)
	}
}
func main() {
	cmd.Execute()
}
