package clazzfile

/*
CONSTANT_MethodHandle {
	u1 tag
	u1 reference_kind
	u2 referenct_index
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (constantMethodHandleInfo *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	constantMethodHandleInfo.referenceKind = reader.readUint8()
	constantMethodHandleInfo.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType {
	u1 tag
	u2 descriptor_index
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (constantMethodTypeInfo *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	constantMethodTypeInfo.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic {
    u1 tag
    u2 bootstrap_method_attr_index
    u2 name_and_type_index
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (constantInvokeDynamicInfo *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	constantInvokeDynamicInfo.bootstrapMethodAttrIndex = reader.readUint16()
	constantInvokeDynamicInfo.nameAndTypeIndex = reader.readUint16()
}
