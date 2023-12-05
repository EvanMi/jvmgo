package control

import (
	"jvmgo/ch10/instructions/base"
	"jvmgo/ch10/rtda"
)

type RETURN struct{ base.NoOperandsInstruction }

func (re *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct{ base.NoOperandsInstruction }

func (aReturn *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokeFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokeFrame.OperandStack().PushRef(ref)
}

type DRETURN struct{ base.NoOperandsInstruction }

func (dReturn *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

type FRETURN struct{ base.NoOperandsInstruction }

func (fReturn *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

type IRETURN struct{ base.NoOperandsInstruction }

func (iReturn *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invlokeFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invlokeFrame.OperandStack().PushInt(val)
}

type LRETURN struct{ base.NoOperandsInstruction }

func (lReturn *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
