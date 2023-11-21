package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type GOTO_W struct {
	offset int
}

func (gotoW *GOTO_W) FetchOperands(reader *base.ByteCodeReader) {
	gotoW.offset = int(reader.ReadInt32())
}

func (gotoW *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, gotoW.offset)
}
