package lang

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

const jlClassLoader = "java/lang/ClassLoader"

func init() {
	native.Register(jlClassLoader, "findBuiltinLib", "(Ljava/lang/String;)Ljava/lang/String;", findBuiltinLib)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func findBuiltinLib(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}
