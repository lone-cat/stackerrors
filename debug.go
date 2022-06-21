package stackerrors

var debugMode bool

const defaultFileName = `"code-point-unavailable"`

func SetDebugMode(mode bool) {
	debugMode = mode
}
