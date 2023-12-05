package control

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (aGoto *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, aGoto.Offset)
}
