package heap

import "jvmgo/ch06/clazzfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *clazzfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (methodRef *MethodRef) ResolvedMethod() *Method {
	if methodRef.method == nil {
		methodRef.resolveMethodRef()
	}
	return methodRef.method
}

func (methodRef *MethodRef) resolveMethodRef() {

}
