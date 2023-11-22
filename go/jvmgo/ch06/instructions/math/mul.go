package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type DMUL struct{ base.NoOperandsInstruction }

func (dmul *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1 * v2
	stack.PushDouble(res)
}

type FMUL struct{ base.NoOperandsInstruction }

func (fmul *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1 * v2
	stack.PushFloat(res)
}

type IMUL struct{ base.NoOperandsInstruction }

func (imul *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1 * v2
	stack.PushInt(res)
}

type LMUL struct{ base.NoOperandsInstruction }

func (lmul *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1 * v2
	stack.PushLong(res)
}
