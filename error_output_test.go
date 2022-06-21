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
	ctxs := [3]string{`root Context`, `package Context`, `func Context`}
	err = stackerrors.Wrap(ctxs[2], err)
	err = stackerrors.Wrap(ctxs[1], err)
	err = stackerrors.Wrap(ctxs[0], err)
	_, file, line, _ := runtime.Caller(0)
	lines := [4]string{``, ``, ``, ``}
	lineNumbers := [3]int{line - 1, line - 2, line - 3}
	for x := 0; x < 3; x++ {
		lines[x] = fmt.Sprintf(`%s %s:%d > `, ctxs[x], file, lineNumbers[x])
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
