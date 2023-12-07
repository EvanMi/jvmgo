package clazzfile

/*
CONSTANT_NameAndType {
    u1 tag
    u2 name_index
    u2 descriptor_index
}
*/

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (constantNameAndTypeInfo *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	constantNameAndTypeInfo.nameIndex = reader.readUint16()
	constantNameAndTypeInfo.descriptorIndex = reader.readUint16()
}
