package stack

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type SWAP struct{ base.NoOperandsInstruction }

func (swap *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
