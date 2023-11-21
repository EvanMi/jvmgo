package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type FSTORE struct{ base.Index8Instruction }

func (fstore *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, fstore.Index)
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (fstore0 *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (fstore1 *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (fstore2 *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (fstore3 *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
