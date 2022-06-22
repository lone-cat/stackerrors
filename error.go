package stackerrors

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
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
	for err != nil {
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
			err = nil
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

func New(text string) error {
	err := errors.New(text)
	return wrap(err, 1)
}

func Newf(format string, args ...any) error {
	err := fmt.Errorf(format, args...)
	return wrap(err, 1)
}

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return wrap(err, 1)
}

func WrapInDefer(err error) error {
	return wrap(err, 2)
}

func wrap(err error, depth uint) error {
	if err == nil {
		return nil
	}

	context := defaultContext
	file := defaultFileName
	var line int

	pc, filename, lineNum, ok := runtime.Caller(int(depth) + 1)

	if ok {
		file = filename
		line = lineNum
	}

	funcName := getShortFuncNameFromPc(pc)
	if funcName != `` {
		context = funcName
	}

	//frame, _ := runtime.CallersFrames([]uintptr{pc}).Next()

	return &DebugContextError{
		context:      context,
		wrappedError: err,
		file:         file,
		line:         line,
	}
}

func getShortFuncNameFromPc(pc uintptr) (funcName string) {
	fun := runtime.FuncForPC(pc)

	if fun == nil {
		return
	}

	parts := strings.Split(fun.Name(), `.`)
	funcName = parts[len(parts)-1]

	if funcName != `` {
		funcName += `()`
	}

	return
}
