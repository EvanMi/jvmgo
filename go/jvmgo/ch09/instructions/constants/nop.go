package constants

import (
	"jvmgo/ch09/instructions/base"
	"jvmgo/ch09/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (nop *NOP) Execute(frame *rtda.Frame) {

}
