package heap

import "jvmgo/ch09/clazzfile"

type ClassRef struct {
	SymRef
}

func newClassRef(cp *ConstantPool, classInfo *clazzfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
