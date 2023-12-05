package constants

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type BIPUSH struct {
	val int8
}

func (bipush *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	bipush.val = reader.ReadInt8()
}

func (bipush *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(bipush.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (sipush *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	sipush.val = reader.ReadInt16()
}

func (sipush *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(sipush.val)
	frame.OperandStack().PushInt(i)
}
