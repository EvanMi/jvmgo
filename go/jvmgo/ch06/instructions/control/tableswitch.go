package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (tableSwitch *TABLE_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	tableSwitch.defaultOffset = reader.ReadInt32()
	tableSwitch.low = reader.ReadInt32()
	tableSwitch.high = reader.ReadInt32()
	jumpOffsetsCount := tableSwitch.high - tableSwitch.low + 1
	tableSwitch.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (tableSwitch *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	var offset int
	if index >= tableSwitch.low && index <= tableSwitch.high {
		offset = int(tableSwitch.jumpOffsets[index])
	} else {
		offset = int(tableSwitch.defaultOffset)
	}
	base.Branch(frame, offset)
}
