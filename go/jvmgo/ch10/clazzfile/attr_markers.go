package clazzfile

type MarkerAttribute struct{}

func (markerAttribute *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}
