package conversions

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type I2B struct{ base.NoOperandsInstruction }

func (i2b *I2B) Execute(frame *rtda.Frame) {
	_i2x(frame, func(i int32, stack *rtda.OperandStack) {
		b := int32(int8(i))
		stack.PushInt(b)
	})
}

type I2C struct{ base.NoOperandsInstruction }

func (i2c *I2C) Execute(frame *rtda.Frame) {
	_i2x(frame, func(i int32, stack *rtda.OperandStack) {
		c := int32(int16(i))
		stack.PushInt(c)
	})
}

type I2S struct{ base.NoOperandsInstruction }

func (i2s *I2S) Execute(frame *rtda.Frame) {
	_i2x(frame, func(i int32, stack *rtda.OperandStack) {
		s := int32(int16(i))
		stack.PushInt(s)
	})
}

type I2L struct{ base.NoOperandsInstruction }

func (i2l *I2L) Execute(frame *rtda.Frame) {
	_i2x(frame, func(i int32, stack *rtda.OperandStack) {
		l := int64(i)
		stack.PushLong(l)
	})
}

type I2F struct{ base.NoOperandsInstruction }

func (i2f *I2F) Execute(frame *rtda.Frame) {
	_i2x(frame, func(i int32, stack *rtda.OperandStack) {
		f := float32(i)
		stack.PushFloat(f)
	})
}

type I2D struct{ base.NoOperandsInstruction }

func (i2d *I2D) Execute(frame *rtda.Frame) {
	_i2x(frame, func(i int32, stack *rtda.OperandStack) {
		d := float64(i)
		stack.PushDouble(d)
	})
}

func _i2x(frame *rtda.Frame, consume func(int32, *rtda.OperandStack)) {
	stack := frame.OperandStack()
	i := stack.PopInt()
	consume(i, stack)
}
