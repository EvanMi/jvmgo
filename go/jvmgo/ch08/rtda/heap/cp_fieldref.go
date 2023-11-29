package heap

import "jvmgo/ch08/clazzfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *clazzfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (fieldRef *FieldRef) ResolvedField() *Field {
	if fieldRef.field == nil {
		fieldRef.resolveFieldRef()
	}
	return fieldRef.field
}

func (fieldRef *FieldRef) resolveFieldRef() {
	d := fieldRef.cp.class
	c := fieldRef.ResolvedClass()
	field := lookupField(c, fieldRef.name, fieldRef.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	fieldRef.field = field
}

func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}

	return nil
}
