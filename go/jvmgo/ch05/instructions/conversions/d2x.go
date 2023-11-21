package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type D2F struct{ base.NoOperandsInstruction }

func (d2f *D2F) Execute(frame *rtda.Frame) {
	d2x(frame, func(d float64, stack *rtda.OperandStack) {
		f := float32(d)
		stack.PushFloat(f)
	})
}

type D2I struct{ base.NoOperandsInstruction }

func (d2i *D2I) Execute(frame *rtda.Frame) {
	d2x(frame, func(d float64, stack *rtda.OperandStack) {
		i := int32(d)
		stack.PushInt(i)
	})
}

type D2L struct{ base.NoOperandsInstruction }

func (d2l *D2L) Execute(frame *rtda.Frame) {
	d2x(frame, func(d float64, stack *rtda.OperandStack) {
		l := int64(d)
		stack.PushLong(l)
	})
}

func d2x(frame *rtda.Frame, f func(float64, *rtda.OperandStack)) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f(d, stack)
}
