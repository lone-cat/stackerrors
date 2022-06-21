package stackerrors

import (
	"errors"
	"fmt"
	"runtime"
)

const defaultSeparator = ` > `

type unwrapper interface {
	Unwrap() error
}

type DebugContextError struct {
	context      string
	wrappedError error
	file         string
	line         int
}

func (e *DebugContextError) Error() string {
	separator := defaultSeparator
	if debugMode {
		separator += "\r\n"
	}
	var err error = e
	var debugSubErr *DebugContextError
	var subErr unwrapper
	var ok bool
	var result string
	var line string
	for {
		if err == nil {
			break
		}
		subErr, ok = err.(unwrapper)
		if ok {
			debugSubErr, ok = subErr.(*DebugContextError)
			if ok {
				line = debugSubErr.context
				if debugMode {
					if line != `` {
						line += ` `
					}
					line += debugSubErr.file
					if debugSubErr.line > 0 {
						line += fmt.Sprintf(`:%d`, debugSubErr.line)
					}
				}
			} else {
				line = err.Error()
			}
			if line != `` {
				result += line + separator
			}
			err = subErr.Unwrap()
		} else {
			result += err.Error()
			break
		}
	}
	return result
}

func (e *DebugContextError) Unwrap() error {
	return e.wrappedError
}

func (e *DebugContextError) File() string {
	return e.file
}

func (e *DebugContextError) Line() int {
	return int(e.line)
}

func New(context string, text string) *DebugContextError {
	err := errors.New(text)
	return wrap(context, err, 1)
}

func Newf(context string, format string, args ...any) *DebugContextError {
	err := fmt.Errorf(format, args...)
	return wrap(context, err, 1)
}

func Wrap(context string, err error) *DebugContextError {
	return wrap(context, err, 1)
}

func WrapInDefer(context string, err error) *DebugContextError {
	return wrap(context, err, 2)
}

func wrap(context string, err error, depth uint) *DebugContextError {
	if err == nil {
		return nil
	}

	file := defaultFileName
	var line int

	if debugMode {
		_, filename, lineNum, ok := runtime.Caller(int(depth) + 1)

		if ok {
			file = filename
			line = lineNum
		}
	}

	//frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()

	return &DebugContextError{
		context:      context,
		wrappedError: err,
		file:         file,
		line:         line,
	}
}
