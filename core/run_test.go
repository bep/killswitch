package core

import (
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

const limit = 4

func TestRun(t *testing.T) {

	testExecutable := "/usr/bin/yes"

	if runtime.GOOS == "windows" {
		testExecutable = "notepad.exe"
	}

	ctx := &Context{Executable: testExecutable, Interval: 1}
	c := &countingConditional{counter: 0, limit: limit}
	var finished uint64

	go func() {
		err := Run(ctx, c)
		if err != nil {
			t.Fatalf("Run() returned err: %s", err)
		}
		atomic.AddUint64(&finished, 1)

	}()

	time.Sleep(time.Second * limit)

	if !atomic.CompareAndSwapUint64(&finished, 1, 2) {
		t.Errorf("Run() did not finish in time")
	} else {
		if c.counter != limit {
			t.Errorf("Counter not reached its limit: %d", c.counter)

		}
	}

}

type countingConditional struct {
	counter int
	limit   int
}

func (c *countingConditional) Valid(ctx *Context) (bool, error) {
	c.counter++
	return c.counter < c.limit, nil
}
