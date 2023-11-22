package heap

import "jvmgo/ch06/clazzfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*clazzfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

func (method *Method) copyAttributes(cfMethod *clazzfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
		method.code = codeAttr.Code()
	}
}

func (method *Method) IsSynchronized() bool {
	return method.accessFlags&ACC_SYNCHRONIZED != 0
}
func (method *Method) IsBridge() bool {
	return method.accessFlags&ACC_BRIDGE != 0
}
func (method *Method) IsVarargs() bool {
	return method.accessFlags&ACC_VARARGS != 0
}
func (method *Method) IsNative() bool {
	return method.accessFlags&ACC_NATIVE != 0
}
func (method *Method) IsAbstract() bool {
	return method.accessFlags&ACC_ABSTRACT != 0
}
func (method *Method) IsStrict() bool {
	return method.accessFlags&ACC_STRICT != 0
}

func (method *Method) MaxStack() uint {
	return method.maxStack
}
func (method *Method) MaxLocals() uint {
	return method.maxLocals
}
func (method *Method) Code() []byte {
	return method.code
}
