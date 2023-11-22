package loads

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

// Load int from local variable
type ILOAD struct{ base.Index8Instruction }

func (iload *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, iload.Index)
}

type ILOAD_0 struct{ base.NoOperandsInstruction }

func (iload0 *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct{ base.NoOperandsInstruction }

func (iload1 *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct{ base.NoOperandsInstruction }

func (iload2 *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct{ base.NoOperandsInstruction }

func (iload3 *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
