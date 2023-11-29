package comparisons

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type IF_ACMPEQ struct{ base.BranchInstruction }

func (ifAcmpeq *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	if _acmp(frame) {
		base.Branch(frame, ifAcmpeq.Offset)
	}
}

type IF_ACMPNE struct{ base.BranchInstruction }

func (ifAcmpne *IF_ACMPNE) Execute(frame *rtda.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, ifAcmpne.Offset)
	}
}

func _acmp(frame *rtda.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2
}
