package stores

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type DSTORE struct{ base.Index8Instruction }

func (destore *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, destore.Index)
}

type DSTORE_0 struct{ base.NoOperandsInstruction }

func (destore0 *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct{ base.NoOperandsInstruction }

func (dstore1 DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct{ base.NoOperandsInstruction }

func (dstore2 DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct{ base.NoOperandsInstruction }

func (destore3 DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
