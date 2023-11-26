package base

type ByteCodeReader struct {
	code []byte
	pc   int
}

func (reader *ByteCodeReader) Reset(code []byte, pc int) {
	reader.code = code
	reader.pc = pc
}

func (reader *ByteCodeReader) PC() int {
	return reader.pc
}

func (reader *ByteCodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *ByteCodeReader) ReadUint8() uint8 {
	i := reader.code[reader.pc]
	reader.pc++
	return i
}

func (reader *ByteCodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *ByteCodeReader) ReadUint16() uint16 {
	byte1 := uint16(reader.ReadUint8())
	byte2 := uint16(reader.ReadUint8())
	return (byte1 << 8) | byte2
}

func (reader *ByteCodeReader) ReadInt32() int32 {
	byte1 := int32(reader.ReadUint8())
	byte2 := int32(reader.ReadUint8())
	byte3 := int32(reader.ReadUint8())
	byte4 := int32(reader.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (reader *ByteCodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = reader.ReadInt32()
	}
	return ints
}

func (reader *ByteCodeReader) SkipPadding() {
	for reader.pc%4 != 0 {
		reader.ReadUint8()
	}
}
