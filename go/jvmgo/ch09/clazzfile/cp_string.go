package clazzfile

/*
CONSTANT_String {
    u1 tag
    u2 string_index
}
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (constantStringInfo *ConstantStringInfo) readInfo(reader *ClassReader) {
	constantStringInfo.stringIndex = reader.readUint16()
}
func (constantStringInfo *ConstantStringInfo) String() string {
	return constantStringInfo.cp.getUtf8(constantStringInfo.stringIndex)
}
