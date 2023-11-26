package control

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (aGoto *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, aGoto.Offset)
}
