package conversions

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtda"
)

type L2D struct{ base.NoOperandsInstruction }

func (l2d *L2D) Execute(frame *rtda.Frame) {
	_l2x(frame, func(l int64, stack *rtda.OperandStack) {
		d := float64(l)
		stack.PushDouble(d)
	})
}

type L2F struct{ base.NoOperandsInstruction }

func (l2f *L2F) Execute(frame *rtda.Frame) {
	_l2x(frame, func(l int64, stack *rtda.OperandStack) {
		f := float32(l)
		stack.PushFloat(f)
	})
}

type L2I struct{ base.NoOperandsInstruction }

func (l2i *L2I) Execute(frame *rtda.Frame) {
	_l2x(frame, func(l int64, stack *rtda.OperandStack) {
		i := int32(l)
		stack.PushInt(i)
	})
}

func _l2x(frame *rtda.Frame, consume func(int64, *rtda.OperandStack)) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	consume(l, stack)
}
