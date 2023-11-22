package heap

import "jvmgo/ch06/clazzfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *clazzfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (interfaceMethodRef *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if interfaceMethodRef.method == nil {
		interfaceMethodRef.resolveInterfaceMethodRef()
	}
	return interfaceMethodRef.method
}

// jvms8 5.4.3.4
func (interfaceMethodRef *InterfaceMethodRef) resolveInterfaceMethodRef() {
	//class := self.ResolveClass()
	// todo
}
