package references

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
	"jvmgo/ch10/rtda/heap"
)

type PUT_STATIC struct{ base.Index16Instruction }

func (putStatic *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(putStatic.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		slots.SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		slots.SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		slots.SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		slots.SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		slots.SetRef(slotId, val)
	default:
		//todo
	}
}
