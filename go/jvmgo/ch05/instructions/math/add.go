package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type DADD struct{ base.NoOperandsInstruction }

func (dadd *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1 + v2
	stack.PushDouble(res)
}

type FADD struct{ base.NoOperandsInstruction }

func (fadd *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1 + v2
	stack.PushFloat(res)
}

type IADD struct{ base.NoOperandsInstruction }

func (iadd *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1 + v2
	stack.PushInt(res)
}

type LADD struct{ base.NoOperandsInstruction }

func (ladd *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1 + v2
	stack.PushLong(res)
}
