package stackerrors

func NilError() *DebugContextError {
	var err *DebugContextError
	return err
}

func GetShortFuncNameFromPc(pc uintptr) string {
	return getShortFuncNameFromPc(pc)
}
