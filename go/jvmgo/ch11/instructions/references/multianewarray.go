package references

import (
	"jvmgo/ch11/instructions/base"
	"jvmgo/ch11/rtda"
	"jvmgo/ch11/rtda/heap"
)

type MULTI_ANEW_ARRAY struct {
	index      uint16
	dimensions uint8
}

func (multiAnewArray *MULTI_ANEW_ARRAY) FetchOperands(reader *base.ByteCodeReader) {
	multiAnewArray.index = reader.ReadUint16()
	multiAnewArray.dimensions = reader.ReadUint8()
}

func (multiAnewArray *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(uint(multiAnewArray.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()

	stack := frame.OperandStack()
	counts := popAndCheckCounts(stack, int(multiAnewArray.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)

}

func popAndCheckCounts(stack *rtda.OperandStack, dimesions int) []int32 {
	counts := make([]int32, dimesions)
	for i := dimesions - 1; i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)
	if len(counts) > 1 {
		refs := arr.Refs()
		for i := range refs {
			refs[i] = newMultiDimensionalArray(counts[1:], arrClass.ComponentClass())
		}
	}
	return arr
}
