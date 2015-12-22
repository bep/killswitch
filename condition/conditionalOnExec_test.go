package condition

import (
	"github.com/bep/killswitch/core"
	"os"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"testing"
	"time"
)

func TestConditionalOnExec(t *testing.T) {

	validationScript := "fail_on_5.sh"
	testExecutable := "/usr/bin/yes"

	if runtime.GOOS == "windows" {
		validationScript = "fail_on_5.cmd"
		testExecutable = "notepad.exe"
	}

	currDir, _ := os.Getwd()
	testFilesDir := filepath.Join(currDir, "..", "testfiles")
	testScript := filepath.Join(testFilesDir, validationScript)

	conditional := newExecConditionalWithCounter(testScript)

	ctx := &core.Context{Executable: testExecutable, Interval: 2}

	var finished uint64

	go func() {
		err := core.Run(ctx, conditional)
		if err != nil {
			t.Fatalf("Run() returned err: %s", err)
		}
		atomic.AddUint64(&finished, 1)

	}()

	time.Sleep(time.Second * 10)

	if !atomic.CompareAndSwapUint64(&finished, 1, 2) {
		t.Errorf("Run() did not finish in time")
	} else {
		if conditional.getOpCounter() != 5 {
			t.Errorf("Counter not reached its limit: %d", conditional.getOpCounter())

		}
	}
}
