package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

type IF_ICMPEQ struct{ base.BranchInstruction }

func (ifIcmpeq *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, ifIcmpeq.Offset)
	}
}

type IF_ICMPNE struct{ base.BranchInstruction }

func (ifIcmpne *IF_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, ifIcmpne.Offset)
	}
}

type IF_ICMPLT struct{ base.BranchInstruction }

func (ifIcmplt *IF_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, ifIcmplt.Offset)
	}
}

type IF_ICMPLE struct{ base.BranchInstruction }

func (ifIcmple *IF_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, ifIcmple.Offset)
	}
}

type IF_ICMPGT struct{ base.BranchInstruction }

func (ifIcmpgt *IF_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, ifIcmpgt.Offset)
	}
}

type IF_ICMPGE struct{ base.BranchInstruction }

func (ifIcmpge *IF_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, ifIcmpge.Offset)
	}
}

func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}
