package math

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type ISHL struct{ base.NoOperandsInstruction }

func (ishl *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	res := v1 << s
	stack.PushInt(res)
}

type ISHR struct{ base.NoOperandsInstruction }

func (ishr *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	res := v1 >> s
	stack.PushInt(res)
}

type IUSHR struct{ base.NoOperandsInstruction }

func (iushr *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	res := int32(uint32(v1) >> s)
	stack.PushInt(res)
}

type LSHL struct{ base.NoOperandsInstruction }

func (lshl *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := v2 & 0x3f
	res := v1 << s
	stack.PushLong(res)
}

type LSHR struct{ base.NoOperandsInstruction }

func (lshr *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := v2 & 0x3f
	res := v1 >> s
	stack.PushLong(res)
}

type LUSHR struct{ base.NoOperandsInstruction }

func (lushr *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	s := v2 & 0x3f
	res := int64(uint64(v1) >> s)
	stack.PushLong(res)
}
