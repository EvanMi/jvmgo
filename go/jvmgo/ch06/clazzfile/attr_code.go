package clazzfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (codeAttr *CodeAttribute) readInfo(reader *ClassReader) {
	codeAttr.maxStack = reader.readUint16()
	codeAttr.maxLocals = reader.readUint16()
	codeLen := reader.readUint32()
	codeAttr.code = reader.readBytes(codeLen)
	codeAttr.exceptionTable = readExceptionTable(reader)
	codeAttr.attributes = readAttributes(reader, codeAttr.cp)
}

func (codeAttr *CodeAttribute) MaxStack() uint {
	return uint(codeAttr.maxStack)
}

func (codeAttr *CodeAttribute) MaxLocals() uint {
	return uint(codeAttr.maxLocals)
}

func (codeAttr *CodeAttribute) Code() []byte {
	return codeAttr.code
}

func (codeAttr *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return codeAttr.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (exceptionEntry *ExceptionTableEntry) StartPc() uint16 {
	return exceptionEntry.startPc
}

func (exceptionEntry *ExceptionTableEntry) EndPc() uint16 {
	return exceptionEntry.endPc
}
func (exceptionEntry *ExceptionTableEntry) HandlerPc() uint16 {
	return exceptionEntry.handlerPc
}
func (exceptionEntry *ExceptionTableEntry) CatchType() uint16 {
	return exceptionEntry.catchType
}
