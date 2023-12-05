package math

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type IXOR struct{ base.NoOperandsInstruction }

func (ixor *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1 ^ v2
	stack.PushInt(res)
}

type LXOR struct{ base.NoOperandsInstruction }

func (lxor *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1 ^ v2
	stack.PushLong(res)
}
