package clazzfile

/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;
        u2 outer_class_info_index;
        u2 inner_name_index;
        u2 inner_class_access_flags;
    } classes[number_of_classes];
}
*/

type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}

type InnerClassInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIdex    uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (innerClassesAttribute *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numOfClasses := reader.readUint16()
	innerClassesAttribute.classes = make([]*InnerClassInfo, numOfClasses)
	for i := range innerClassesAttribute.classes {
		innerClassesAttribute.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIdex:    reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}
