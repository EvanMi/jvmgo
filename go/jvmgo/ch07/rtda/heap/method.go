package heap

import "jvmgo/ch07/clazzfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*clazzfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
		methods[i].calcArgSlotCount()
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

func (method *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(method.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		method.argSlotCount++
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}
	if !method.IsStatic() {
		method.argSlotCount++
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
func (method *Method) ArgSlotCount() uint {
	return method.argSlotCount
}
