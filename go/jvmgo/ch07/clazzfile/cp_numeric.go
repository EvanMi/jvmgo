package clazzfile

import "math"

/*
CONSTANT_Integer {
    u1 tag
    u4 bytes
}
*/

type ConstantIntegerInfo struct {
	val int32
}

func (constantIntegerInfo *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	constantIntegerInfo.val = int32(bytes)
}
func (constantIntegerInfo *ConstantIntegerInfo) Value() int32 {
	return constantIntegerInfo.val
}

/*
CONSTANT_Float {
    u1 tag
    u4 bytes
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (constantFloatInfo *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	constantFloatInfo.val = math.Float32frombits(bytes)
}
func (constantFloatInfo *ConstantFloatInfo) Value() float32 {
	return constantFloatInfo.val
}

/*
CONSTANT_Long {
    u1 tag
    u4 high_bytes
    u4 low_bytes
}
*/
type ConstantLongInfo struct {
	val int64
}

func (constantLongInfo *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	constantLongInfo.val = int64(bytes)
}
func (constantLongInfo *ConstantLongInfo) Value() int64 {
	return constantLongInfo.val
}

/*
CONSTANT_Double {
    u1 tag
    u4 high_bytes
    u4 low_bytes
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (constantDoubleInfo *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	constantDoubleInfo.val = math.Float64frombits(bytes)
}
func (constantDoubleInfo *ConstantDoubleInfo) Value() float64 {
	return constantDoubleInfo.val
}
