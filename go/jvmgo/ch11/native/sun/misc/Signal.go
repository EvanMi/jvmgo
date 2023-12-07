package misc

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

func init() {
	_signal(findSignal, "findSignal", "(Ljava/lang/String;)I")
	_signal(handle0, "handle0", "(IJ)J")
}

func _signal(method func(frame *rtda.Frame), name, desc string) {
	native.Register("sun/misc/Signal", name, desc, method)
}

func findSignal(frame *rtda.Frame) {
	vars := frame.LocalVars()
	vars.GetRef(0)
	stack := frame.OperandStack()
	stack.PushInt(0)
}

func handle0(frame *rtda.Frame) {
	vars := frame.LocalVars()
	vars.GetInt(0)
	vars.GetLong(1)

	stack := frame.OperandStack()
	stack.PushLong(0)
}
