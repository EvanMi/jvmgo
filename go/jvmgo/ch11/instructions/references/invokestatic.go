package references

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

type INVOKE_STATIC struct{ base.Index16Instruction }

func (invokeStatic *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(invokeStatic.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	methodRef.ResolvedClass()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	base.InvokeMethod(frame, resolvedMethod)
}
