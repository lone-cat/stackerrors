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

var inplaceExpectedLine int

func TestErrorInPlace(t *testing.T) {
	err := inPlaceErrorGenerator()
	wrappedErr := stackerrors.NilError()
	ok := errors.As(err, &wrappedErr)
	if !ok {
		t.Fail()
	}
	if wrappedErr.Line() != inplaceExpectedLine {
		fmt.Printf("expected inplace error line is %d, actual is %d\r\n", inplaceExpectedLine, wrappedErr.Line())
		t.Fail()
	}
}

func inPlaceErrorGenerator() (err error) {
	err = defaultErr
	err = stackerrors.Wrap(``, err)
	_, _, inplaceExpectedLine, _ = runtime.Caller(0)
	inplaceExpectedLine--
	return
}
