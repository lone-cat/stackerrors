package stackerrors_test

import (
	"errors"
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
		t.Error(`err == defaultErr`)
	}
}

func TestUnwrap(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)

	wrappedErr := stackerrors.NilError()
	errors.As(err, &wrappedErr)
	if wrappedErr.Unwrap() != defaultErr {
		t.Error(`err.Unwrap() != defaultErr`)
	}
}

func TestIs(t *testing.T) {
	err := defaultErr
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	if !errors.Is(err, defaultErr) {
		t.Error(`!errors.Is(err, defaultErr)`)
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
		t.Error(`errors.As(err, wrappedErr) fails`)
	}
}

func TestNilWrap(t *testing.T) {
	var err error = nil
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	err = stackerrors.Wrap(err)
	if err != nil {
		t.Error(`wrapped nil is stackerror`)
	}
}
