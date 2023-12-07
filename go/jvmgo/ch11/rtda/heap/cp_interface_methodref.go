package heap

import "jvmgo/ch11/clazzfile"

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
	d := interfaceMethodRef.cp.class
	c := interfaceMethodRef.ResolvedClass()
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, interfaceMethodRef.name, interfaceMethodRef.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	interfaceMethodRef.method = method
}

func lookupInterfaceMethod(iface *Class, name, descriptor string) *Method {
	ifaces := make([]*Class, 1)
	ifaces[0] = iface
	return lookupMethodInInterfaces(ifaces, name, descriptor)
}
