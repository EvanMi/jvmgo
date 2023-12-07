package io

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
