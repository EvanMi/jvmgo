package references

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

func (arrayLength *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
