package math

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type DDIV struct{ base.NoOperandsInstruction }

func (ddiv *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	res := v1 / v2
	stack.PushDouble(res)
}

type FDIV struct{ base.NoOperandsInstruction }

func (fdiv *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	res := v1 / v2
	stack.PushFloat(res)
}

type IDIV struct{ base.NoOperandsInstruction }

func (idiv *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := v1 / v2
	stack.PushInt(res)
}

type LDIV struct{ base.NoOperandsInstruction }

func (ldiv *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	res := v2 / v1
	stack.PushLong(res)
}
