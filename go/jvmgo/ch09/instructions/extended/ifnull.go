package extended

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type IFNULL struct{ base.BranchInstruction }

func (ifNull *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, ifNull.Offset)
	}
}

type IFNONNULL struct{ base.BranchInstruction }

func (ifNonNull *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, ifNonNull.Offset)
	}
}
