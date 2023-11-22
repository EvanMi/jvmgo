package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (lookupSwitch *LOOKUP_SWITCH) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	lookupSwitch.defaultOffset = reader.ReadInt32()
	lookupSwitch.npairs = reader.ReadInt32()
	lookupSwitch.matchOffsets = reader.ReadInt32s(lookupSwitch.npairs * 2)
}

func (lookupSwitch *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < lookupSwitch.npairs*2; i += 2 {
		if lookupSwitch.matchOffsets[i] == key {
			offset := lookupSwitch.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(lookupSwitch.defaultOffset))
}
