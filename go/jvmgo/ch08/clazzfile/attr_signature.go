package clazzfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (signatureAttribute *SignatureAttribute) readInfo(reader *ClassReader) {
	signatureAttribute.signatureIndex = reader.readUint16()
}

func (signatureAttribute *SignatureAttribute) Signature() string {
	return signatureAttribute.cp.getUtf8(signatureAttribute.signatureIndex)
}
