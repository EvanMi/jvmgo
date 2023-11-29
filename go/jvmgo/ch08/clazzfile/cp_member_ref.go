package clazzfile

/*
CONSTANT_Fieldref {
    u1 tag
    u2 class_index
    u2 name_and_type_index
}
CONSTANT_Methodref {
    u1 tag
    u2 class_index
    u2 name_and_type_index
}
CONSTANT_InterfaceMethodref {
    u1 tag
    u2 class_index
    u2 name_and_type_index
}
*/

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (constantMemberrefInfo *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	constantMemberrefInfo.classIndex = reader.readUint16()
	constantMemberrefInfo.nameAndTypeIndex = reader.readUint16()
}

func (constantMemberrefInfo *ConstantMemberrefInfo) ClassName() string {
	return constantMemberrefInfo.cp.getClassName(constantMemberrefInfo.classIndex)
}

func (constantMemberrefInfo *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return constantMemberrefInfo.cp.getNameAndType(constantMemberrefInfo.nameAndTypeIndex)
}

type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
