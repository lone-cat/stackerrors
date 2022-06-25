package stackerrors

import (
	"errors"
	"io"
	"runtime"
	"testing"
)

func init() {
	SetDebugMode(true)
	defaultErr = io.EOF
}

var deferExpectedLine int

func TestErrorInDefer(t *testing.T) {
	err := inDeferErrorGenerator()
	var wrappedErr *DebugContextError
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
			err = WrapInDefer(err)
		}
	}()

	//wrappedError = defaultErr

	_, _, deferExpectedLine, _ = runtime.Caller(0)
	deferExpectedLine += 2
	return defaultErr
}
