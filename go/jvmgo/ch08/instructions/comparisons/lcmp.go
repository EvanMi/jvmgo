package comparisons

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type LCMP struct{ base.NoOperandsInstruction }

func (lcmp *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
