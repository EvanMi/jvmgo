package constants

import (
	"jvmgo/ch07/instructions/base"
	"jvmgo/ch07/rtda"
)

type LDC struct{ base.Index8Instruction }

func (ldc *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, ldc.Index)
}

type LDC_W struct{ base.Index16Instruction }

func (ldcw *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, ldcw.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c := c.(type) {
	case int32:
		stack.PushInt(c)
	case float32:
		stack.PushFloat(c)
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

type LDC2_W struct{ base.Index16Instruction }

func (ldc2w *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(ldc2w.Index)

	switch c := c.(type) {
	case int64:
		stack.PushLong(c)
	case float64:
		stack.PushDouble(c)
	default:
		panic("java.lang.ClassFormatError")
	}
}
