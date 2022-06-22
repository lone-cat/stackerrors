package stackerrors

var debugMode bool

const defaultContext = `"context-unavailable"`
const defaultFileName = `"code-point-unavailable"`

func SetDebugMode(mode bool) {
	debugMode = mode
}
