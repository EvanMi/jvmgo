package conversions

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type F2D struct{ base.NoOperandsInstruction }

func (f2d *F2D) Execute(frame *rtda.Frame) {
	_f2x(frame, func(f float32, stack *rtda.OperandStack) {
		d := float64(f)
		stack.PushDouble(d)
	})
}

type F2I struct{ base.NoOperandsInstruction }

func (f2i *F2I) Execute(frame *rtda.Frame) {
	_f2x(frame, func(f float32, stack *rtda.OperandStack) {
		i := int32(f)
		stack.PushInt(i)
	})
}

type F2L struct{ base.NoOperandsInstruction }

func (f2l *F2L) Execute(frame *rtda.Frame) {
	_f2x(frame, func(f float32, stack *rtda.OperandStack) {
		l := int64(f)
		stack.PushLong(l)
	})
}

func _f2x(frame *rtda.Frame, consume func(float32, *rtda.OperandStack)) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	consume(f, stack)
}
