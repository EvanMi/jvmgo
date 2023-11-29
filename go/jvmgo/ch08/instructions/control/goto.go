package control

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type GOTO struct{ base.BranchInstruction }

func (aGoto *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, aGoto.Offset)
}
