package clazzfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (memberInfo *MemberInfo) AccessFlags() uint16 {
	return memberInfo.accessFlags
}

func (memeberInfo *MemberInfo) Name() string {
	return memeberInfo.cp.getUtf8(memeberInfo.nameIndex)
}

func (memberInfo *MemberInfo) Descriptor() string {
	return memberInfo.cp.getUtf8(memberInfo.descriptorIndex)
}

func (memberInfo *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range memberInfo.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (memberInfo *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range memberInfo.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (memberInfo *MemberInfo) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range memberInfo.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (memberInfo *MemberInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return memberInfo.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (memberInfo *MemberInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return memberInfo.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (memberInfo *MemberInfo) AnnotationDefaultAttributeData() []byte {
	return memberInfo.getUnparsedAttributeData("AnnotationDefault")
}

func (memberInfo *MemberInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range memberInfo.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}
