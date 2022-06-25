package stackerrors

import (
	"errors"
	"io"
	"testing"
)

func init() {
	SetDebugMode(true)
	defaultErr = io.EOF
}

func TestWrap(t *testing.T) {
	err := defaultErr
	err = Wrap(err)
	if err == defaultErr {
		t.Error(`err == defaultErr`)
	}
}

func TestUnwrap(t *testing.T) {
	err := defaultErr
	err = Wrap(err)

	var wrappedErr *DebugContextError
	errors.As(err, &wrappedErr)
	if wrappedErr.Unwrap() != defaultErr {
		t.Error(`err.Unwrap() != defaultErr`)
	}
}

func TestIs(t *testing.T) {
	err := defaultErr
	err = Wrap(err)
	err = Wrap(err)
	err = Wrap(err)
	if !errors.Is(err, defaultErr) {
		t.Error(`!errors.Is(err, defaultErr)`)
	}
}

func TestAs(t *testing.T) {
	err := defaultErr
	err = Wrap(err)
	err = Wrap(err)
	err = Wrap(err)

	var wrappedErr *DebugContextError
	ok := errors.As(err, &wrappedErr)
	if !ok {
		t.Error(`errors.As(err, wrappedErr) fails`)
	}
}

func TestNilWrap(t *testing.T) {
	var err error = nil
	err = Wrap(err)
	err = Wrap(err)
	err = Wrap(err)
	if err != nil {
		t.Error(`wrapped nil is stackerror`)
	}
}
