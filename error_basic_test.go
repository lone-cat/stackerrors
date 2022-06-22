package stackerrors_test

import (
	"errors"
	"fmt"
	"github.com/lone-cat/stackerrors"
	"io"
	"testing"
)

func init() {
	stackerrors.SetDebugMode(true)
	defaultErr = io.EOF
}

func TestWrap(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)
	if err == defaultErr {
		fmt.Println(`err == defaultErr`)
		t.Fail()
	}
}

func TestUnwrap(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)

	wrappedErr := stackerrors.NilError()
	errors.As(err, &wrappedErr)
	if wrappedErr.Unwrap() != defaultErr {
		fmt.Println(`err.Unwrap() != defaultErr`)
		t.Fail()
	}
}

func TestIs(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	if !errors.Is(err, defaultErr) {
		fmt.Println(`!errors.Is(err, defaultErr)`)
		t.Fail()
	}
}

func TestAs(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)

	wrappedErr := stackerrors.NilError()
	ok := errors.As(err, &wrappedErr)
	if !ok {
		fmt.Println(`errors.As(err, wrappedErr) fails`)
		t.Fail()
	}
}

func TestNilWrap(t *testing.T) {
	var err error = nil
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	if err != nil {
		fmt.Println(`wrapped nil is stackerror`)
		t.Fail()
	}
}
