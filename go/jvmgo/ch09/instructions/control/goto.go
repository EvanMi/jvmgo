package control

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (aGoto *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, aGoto.Offset)
}
