package rtda

import (
	"jvmgo/ch11/rtda/heap"
	"math"
)

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (localVars LocalVars) SetInt(index uint, val int32) {
	localVars[index].num = val
}

func (loccalVars LocalVars) GetInt(index uint) int32 {
	return loccalVars[index].num
}

func (localVars LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	localVars[index].num = int32(bits)
}

func (localVars LocalVars) GetFloat(index uint) float32 {
	bits := uint32(localVars[index].num)
	return math.Float32frombits(bits)
}

func (localVars LocalVars) SetLong(index uint, val int64) {
	localVars[index].num = int32(val)
	localVars[index+1].num = int32(val >> 32)
}

func (localVars LocalVars) GetLong(index uint) int64 {
	low := uint32(localVars[index].num)
	high := uint32(localVars[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (localVars LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	localVars.SetLong(index, int64(bits))
}

func (localVars LocalVars) GetDouble(index uint) float64 {
	bits := uint64(localVars.GetLong(index))
	return math.Float64frombits(bits)
}

func (localVars LocalVars) SetRef(index uint, ref *heap.Object) {
	localVars[index].ref = ref
}

func (localVars LocalVars) GetRef(index uint) *heap.Object {
	return localVars[index].ref
}

func (localVars LocalVars) SetSlot(index uint, slot Slot) {
	localVars[index] = slot
}

func (localVars LocalVars) GetThis() *heap.Object {
	return localVars.GetRef(0)
}
func (localVars LocalVars) GetBoolean(index uint) bool {
	return localVars.GetInt(index) == 1
}
