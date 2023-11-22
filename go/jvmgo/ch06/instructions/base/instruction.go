package base

import "jvmgo/ch06/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct{}

func (noOperandsInstruction *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (branchInstruction *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	branchInstruction.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (index8Instruction *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	index8Instruction.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (index16Instruction *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	index16Instruction.Index = uint(reader.ReadUint16())
}
