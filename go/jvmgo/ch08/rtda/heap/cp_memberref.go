package heap

import "jvmgo/ch08/clazzfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (memberRef *MemberRef) copyMemberRefInfo(refInfo *clazzfile.ConstantMemberrefInfo) {
	memberRef.className = refInfo.ClassName()
	memberRef.name, memberRef.descriptor = refInfo.NameAndDescriptor()
}

func (memberRef *MemberRef) Name() string {
	return memberRef.name
}
func (memberRef *MemberRef) Descriptor() string {
	return memberRef.descriptor
}
