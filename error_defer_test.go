package stackerrors_test

import (
	"errors"
	"fmt"
	"github.com/lone-cat/stackerrors"
	"io"
	"runtime"
	"testing"
)

func init() {
	stackerrors.SetDebugMode(true)
	defaultErr = io.EOF
}

var deferExpectedLine int

func TestErrorInDefer(t *testing.T) {
	err := inDeferErrorGenerator()
	wrappedErr := stackerrors.NilError()
	ok := errors.As(err, &wrappedErr)
	if !ok {
		t.Fail()
	}
	if wrappedErr.Line() != deferExpectedLine {
		fmt.Printf("expected inplace error line is %d, actual is %d\r\n", deferExpectedLine, wrappedErr.Line())
		t.Fail()
	}
}

func inDeferErrorGenerator() (err error) {
	defer func() {
		if err != nil {
			err = stackerrors.WrapInDefer(``, err)
		}
	}()

	//wrappedError = defaultErr

	_, _, deferExpectedLine, _ = runtime.Caller(0)
	deferExpectedLine += 2
	return defaultErr
}
