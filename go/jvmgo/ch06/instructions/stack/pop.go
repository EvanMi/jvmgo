package stack

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type POP struct{ base.NoOperandsInstruction }

func (pop *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct{ base.NoOperandsInstruction }

func (pop2 *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
