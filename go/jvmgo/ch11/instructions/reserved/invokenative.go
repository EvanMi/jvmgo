package reserved

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"

	_ "jvmgo/ch11/native/java/io"

	_ "jvmgo/ch11/native/java/lang"

	_ "jvmgo/ch11/native/java/security"

	_ "jvmgo/ch11/native/java/util/concurrent/atomic"

	_ "jvmgo/ch11/native/sun/io"

	_ "jvmgo/ch11/native/sun/misc"

	_ "jvmgo/ch11/native/sun/reflect"
)

type INVOKE_NATIVE struct{ base.NoOperandsInstruction }

func (invokeNative *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}
