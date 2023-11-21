package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}

func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

func (frame *Frame) Thread() *Thread {
	return frame.thread
}

func (frame *Frame) NextPC() int {
	return frame.nextPC
}

func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}
