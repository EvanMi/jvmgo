package clazzfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/

type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (unparsedAttribute *UnparsedAttribute) readInfo(reader *ClassReader) {
	unparsedAttribute.info = reader.readBytes(unparsedAttribute.length)
}

func (unparsedAttribute *UnparsedAttribute) Info() []byte {
	return unparsedAttribute.info
}
