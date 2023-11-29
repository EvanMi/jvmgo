package clazzfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (constantValueAttribute *ConstantValueAttribute) readInfo(reader *ClassReader) {
	constantValueAttribute.constantValueIndex = reader.readUint16()
}

func (constantValueAttribute *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return constantValueAttribute.constantValueIndex
}
