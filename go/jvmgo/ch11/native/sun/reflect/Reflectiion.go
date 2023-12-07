package reflect

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

func init() {
	native.Register("sun/reflect/Reflection", "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
	native.Register("sun/reflect/Reflection", "getClassAccessFlags", "(Ljava/lang/Class;)I", getClassAccessFlags)

}

func getCallerClass(frame *rtda.Frame) {
	callerFrame := frame.Thread().GetFrames()[2]
	callerClass := callerFrame.Method().Class().JClass()
	frame.OperandStack().PushRef(callerClass)
}

func getClassAccessFlags(frame *rtda.Frame) {
	vars := frame.LocalVars()
	_type := vars.GetRef(0)

	goClass := _type.Extra().(*heap.Class)
	flags := goClass.AccessFlags()

	stack := frame.OperandStack()
	stack.PushInt(int32(flags))
}
