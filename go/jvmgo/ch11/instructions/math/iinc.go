package math

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (iinc *IINC) FetchOperands(reader *base.ByteCodeReader) {
	iinc.Index = uint(reader.ReadUint8())
	iinc.Const = int32(reader.ReadInt8())
}

func (iinc *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(iinc.Index)
	val += iinc.Const
	localVars.SetInt(iinc.Index, val)
}
