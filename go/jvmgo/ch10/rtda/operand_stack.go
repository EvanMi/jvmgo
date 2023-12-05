package rtda

import (
	"jvmgo/ch10/rtda/heap"
	"math"
)

type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (operandStack *OperandStack) PushInt(val int32) {
	operandStack.slots[operandStack.size].num = val
	operandStack.size++
}

func (operandStack *OperandStack) PopInt() int32 {
	operandStack.size--
	return operandStack.slots[operandStack.size].num
}

func (operandStack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	operandStack.PushInt(int32(bits))
}

func (operandStack *OperandStack) PopFloat() float32 {
	bits := uint32(operandStack.PopInt())
	return math.Float32frombits(bits)
}

func (operandStack *OperandStack) PushLong(val int64) {
	operandStack.slots[operandStack.size].num = int32(val)
	operandStack.slots[operandStack.size+1].num = int32(val >> 32)
	operandStack.size += 2
}

func (operandStack *OperandStack) PopLong() int64 {
	operandStack.size -= 2
	low := uint32(operandStack.slots[operandStack.size].num)
	high := uint32(operandStack.slots[operandStack.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (operandStack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	operandStack.PushLong(int64(bits))
}

func (operandStack *OperandStack) PopDouble() float64 {
	bits := uint64(operandStack.PopLong())
	return math.Float64frombits(bits)
}

func (operandStack *OperandStack) PushRef(ref *heap.Object) {
	operandStack.slots[operandStack.size].ref = ref
	operandStack.size++
}

func (operandStack *OperandStack) PopRef() *heap.Object {
	operandStack.size--
	ref := operandStack.slots[operandStack.size].ref
	//帮助垃圾回收
	operandStack.slots[operandStack.size].ref = nil
	return ref
}

func (operandStack *OperandStack) PushSlot(slot Slot) {
	operandStack.slots[operandStack.size] = slot
	operandStack.size++
}

func (operandStack *OperandStack) PopSlot() Slot {
	operandStack.size--
	return operandStack.slots[operandStack.size]
}

func (operandStack *OperandStack) GetRefFromTop(n uint) *heap.Object {
	return operandStack.slots[operandStack.size-1-n].ref
}

func (operandStack *OperandStack) PushBoolean(val bool) {
	if val {
		operandStack.PushInt(1)
	} else {
		operandStack.PushInt(0)
	}
}
func (operandStack *OperandStack) PopBoolean() bool {
	return operandStack.PopInt() == 1
}
func (operandStack *OperandStack) Clear() {
	operandStack.size = 0
	for i := range operandStack.slots {
		operandStack.slots[i].ref = nil
	}
}
