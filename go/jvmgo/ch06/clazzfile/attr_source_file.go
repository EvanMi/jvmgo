package clazzfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (sourceFileAttribute *SourceFileAttribute) readInfo(reader *ClassReader) {
	sourceFileAttribute.sourceFileIndex = reader.readUint16()
}

func (sourceFileAttribute *SourceFileAttribute) FileName() string {
	return sourceFileAttribute.cp.getUtf8(sourceFileAttribute.sourceFileIndex)
}
