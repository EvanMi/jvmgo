package clazzfile

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	cp         ConstantPool
	classIndex uint16
	methodInex uint16
}

func (enclosingMethodAttribute *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	enclosingMethodAttribute.classIndex = reader.readUint16()
	enclosingMethodAttribute.methodInex = reader.readUint16()
}

func (enclosingMethodAttribute *EnclosingMethodAttribute) ClassName() string {
	return enclosingMethodAttribute.cp.getUtf8(enclosingMethodAttribute.classIndex)
}

func (enclosingMethodAttribute *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if enclosingMethodAttribute.methodInex > 0 {
		return enclosingMethodAttribute.cp.getNameAndType(enclosingMethodAttribute.methodInex)
	}

	return "", ""
}
