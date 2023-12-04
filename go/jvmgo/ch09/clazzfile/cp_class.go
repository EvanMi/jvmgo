package clazzfile

/*
CONSTANT_Class {
	u1 tag
	u2 name_index
}
*/

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (classInfo *ConstantClassInfo) readInfo(reader *ClassReader) {
	classInfo.nameIndex = reader.readUint16()
}

func (classInfo *ConstantClassInfo) Name() string {
	return classInfo.cp.getUtf8(classInfo.nameIndex)
}
