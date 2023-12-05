package loads

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type DLOAD struct{ base.Index8Instruction }

func (dload *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, dload.Index)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (dload0 *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (dload1 *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (dload2 *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (dload3 *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
