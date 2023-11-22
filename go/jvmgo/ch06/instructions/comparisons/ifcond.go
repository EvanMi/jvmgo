package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type IFEQ struct{ base.BranchInstruction }

func (ifeq *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, ifeq.Offset)
	}
}

type IFNE struct{ base.BranchInstruction }

func (ifne *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, ifne.Offset)
	}
}

type IFLT struct{ base.BranchInstruction }

func (iflt *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, iflt.Offset)
	}
}

type IFLE struct{ base.BranchInstruction }

func (ifle *IFLE) Execute(frame *rtda.Frame) {
	if val := frame.OperandStack().PopInt(); val <= 0 {
		base.Branch(frame, ifle.Offset)
	}
}

type IFGT struct{ base.BranchInstruction }

func (ifgt *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, ifgt.Offset)
	}
}

type IFGE struct{ base.BranchInstruction }

func (ifge *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, ifge.Offset)
	}
}
