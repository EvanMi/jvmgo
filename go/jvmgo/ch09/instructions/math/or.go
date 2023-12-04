package math

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type IOR struct{ base.NoOperandsInstruction }

func (ior *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1 | v2
	stack.PushInt(res)
}

type LOR struct{ base.NoOperandsInstruction }

func (lor *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1 | v2
	stack.PushLong(res)
}
