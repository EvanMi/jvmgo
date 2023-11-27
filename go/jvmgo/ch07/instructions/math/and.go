package math

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type IAND struct{ base.NoOperandsInstruction }

func (iadd *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	res := v1 & v2
	stack.PushInt(res)
}

type LAND struct{ base.NoOperandsInstruction }

func (land *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	res := v1 & v2
	stack.PushLong(res)
}
