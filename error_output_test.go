package stackerrors_test

import (
	"fmt"
	"github.com/lone-cat/stackerrors"
	"io"
	"runtime"
	"strings"
	"testing"
)

var defaultErr error

func init() {
	stackerrors.SetDebugMode(true)
	defaultErr = io.EOF
}

func TestOutput(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	pc, file, line, _ := runtime.Caller(0)
	functionName := stackerrors.GetShortFuncNameFromPc(pc)

	lines := [4]string{``, ``, ``, ``}
	lineNumbers := [3]int{line - 1, line - 2, line - 3}
	for x := 0; x < 3; x++ {
		lines[x] = fmt.Sprintf(`%s %s:%d > `, functionName, file, lineNumbers[x])
	}
	lines[3] = defaultErr.Error()
	expectedOut := strings.Join(lines[:], "\r\n")
	out := err.Error()
	if out != expectedOut {
		fmt.Println(out)
		fmt.Println(expectedOut)
		t.Fail()
	}
}
