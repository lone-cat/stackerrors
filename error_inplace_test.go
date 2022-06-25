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

var inplaceExpectedLine int

func TestErrorInPlace(t *testing.T) {
	err := inPlaceErrorGenerator()
	var wrappedErr *DebugContextError
	ok := errors.As(err, &wrappedErr)
	if !ok {
		t.Fatal(`unable to convert error interface to wrappedErr`)
	}
	if wrappedErr.Line() != inplaceExpectedLine {
		t.Errorf(`expected inplace error line is "%d", actual is "%d"`, inplaceExpectedLine, wrappedErr.Line())
	}
}

func inPlaceErrorGenerator() (err error) {
	err = defaultErr
	err = Wrap(err)
	_, _, inplaceExpectedLine, _ = runtime.Caller(0)
	inplaceExpectedLine--
	return
}
