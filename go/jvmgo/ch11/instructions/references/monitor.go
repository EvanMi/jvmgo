package references

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
)

type MONITOR_ENTER struct{ base.NoOperandsInstruction }

func (monitorenter *MONITOR_ENTER) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

// Exit monitor for object
type MONITOR_EXIT struct{ base.NoOperandsInstruction }

// todo
func (monitorexit *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
