package base

import "jvmgo/ch09/rtda"

func Branch(frame *rtda.Frame, Offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + Offset
	frame.SetNextPC(nextPC)
}
