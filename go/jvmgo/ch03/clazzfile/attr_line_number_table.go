package clazzfile

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (lineNumberTableAttribute *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	lineNumberTableAttribute.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range lineNumberTableAttribute.lineNumberTable {
		lineNumberTableAttribute.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

func (lineNumberTableAttribute *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(lineNumberTableAttribute.lineNumberTable) - 1; i >= 0; i-- {
		entry := lineNumberTableAttribute.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
