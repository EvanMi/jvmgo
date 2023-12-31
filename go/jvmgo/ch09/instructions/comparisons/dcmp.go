package comparisons

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type DCMPG struct{ base.NoOperandsInstruction }

func (dcmpg *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct{ base.NoOperandsInstruction }

func (dcmpl *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	// 由于在比较的过程中可能会存在Nan，造成无法比较情况，所以需要在无法比较的时候做出选择
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()

	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
