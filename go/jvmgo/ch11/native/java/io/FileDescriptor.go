package io

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

const fd = "java/io/FileDescriptor"

func init() {
	native.Register(fd, "set", "(I)J", set)
	native.Register(fd, "initIDs", "()V", initIDs)
}

// private static native long set(int d);
// (I)J
func set(frame *rtda.Frame) {
	// todo
	frame.OperandStack().PushLong(0)
}
