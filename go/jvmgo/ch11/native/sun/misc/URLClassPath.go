package misc

import (
	"jvmgo/ch11/native"
	"jvmgo/ch11/rtda"
)

func init() {
	native.Register("sun/misc/URLClassPath", "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

func getLookupCacheURLs(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}
