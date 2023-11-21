package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (nop *NOP) Execute(frame *rtda.Frame) {

}
