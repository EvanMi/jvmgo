package heap

import "jvmgo/ch11/clazzfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc   int
	endPc     int
	handlePc  int
	catchType *ClassRef
}

func newExceptionTable(entries []*clazzfile.ExceptionTableEntry, cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc:   int(entry.StartPc()),
			endPc:     int(entry.EndPc()),
			handlePc:  int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}

func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil //catch all
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (table ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {
	for _, handler := range table {
		// jvms: The start_pc is inclusive and end_pc is exclusive
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}

	}
	return nil
}
