package clazzfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

//u1
func (classReader *ClassReader) readUint8() uint8 {
	val := classReader.data[0]
	classReader.data = classReader.data[1:]
	return val
}

//u2
func (classReader *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(classReader.data)
	classReader.data = classReader.data[2:]
	return val
}

//u4
func (classReader *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(classReader.data)
	classReader.data = classReader.data[4:]
	return val
}

func (classReader *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(classReader.data)
	classReader.data = classReader.data[8:]
	return val
}

func (classReader *ClassReader) readUint16s() []uint16 {
	n := classReader.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = classReader.readUint16()
	}
	return s
}

func (classReader *ClassReader) readBytes(n uint32) []byte {
	bytes := classReader.data[:n]
	classReader.data = classReader.data[n:]
	return bytes
}
