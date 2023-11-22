package heap

import "jvmgo/ch06/clazzfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (classMember *ClassMember) copyMemberInfo(memberInfo *clazzfile.MemberInfo) {
	classMember.accessFlags = memberInfo.AccessFlags()
	classMember.name = memberInfo.Name()
	classMember.descriptor = memberInfo.Descriptor()
}

func (classMember *ClassMember) IsPublic() bool {
	return classMember.accessFlags&ACC_PUBLIC != 0
}
func (classMember *ClassMember) IsPrivate() bool {
	return classMember.accessFlags&ACC_PRIVATE != 0
}
func (classMember *ClassMember) IsProtected() bool {
	return classMember.accessFlags&ACC_PROTECTED != 0
}
func (classMember *ClassMember) IsStatic() bool {
	return classMember.accessFlags&ACC_STATIC != 0
}
func (classMember *ClassMember) IsFinal() bool {
	return classMember.accessFlags&ACC_FINAL != 0
}
func (classMember *ClassMember) IsSynthetic() bool {
	return classMember.accessFlags&ACC_SYNTHETIC != 0
}

func (classMember *ClassMember) Name() string {
	return classMember.name
}
func (classMember *ClassMember) Descriptor() string {
	return classMember.descriptor
}
func (classMember *ClassMember) Class() *Class {
	return classMember.class
}

func (classMember *ClassMember) isAccessibleTo(d *Class) bool {
	if classMember.IsPublic() {
		return true
	}
	c := classMember.class
	if classMember.IsProtected() {
		return d == c || d.isSubClassOf(c) || c.getPackageName() == d.getPackageName()
	}
	if !classMember.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	return d == c
}
