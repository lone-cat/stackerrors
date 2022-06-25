package stackerrors_test

import (
	"errors"
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
		t.Fatal(`unable to convert error interface to wrappedErr`)
	}
	if wrappedErr.Line() != deferExpectedLine {
		t.Errorf(`expected defer error line is "%d", actual is "%d"`, deferExpectedLine, wrappedErr.Line())
	}
}

func inDeferErrorGenerator() (err error) {
	defer func() {
		if err != nil {
			err = stackerrors.WrapInDefer(err)
		}
	}()

	//wrappedError = defaultErr

	_, _, deferExpectedLine, _ = runtime.Caller(0)
	deferExpectedLine += 2
	return defaultErr
}
