package heap

import "jvmgo/ch10/clazzfile"

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            []byte
	exceptionTable  ExceptionTable
	lineNumberTable *clazzfile.LineNumberTableAttribute
	argSlotCount    uint
}

func newMethods(class *Class, cfMethods []*clazzfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *clazzfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (method *Method) copyAttributes(cfMethod *clazzfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		method.maxStack = codeAttr.MaxStack()
		method.maxLocals = codeAttr.MaxLocals()
		method.code = codeAttr.Code()
		method.lineNumberTable = codeAttr.LineNumberTableAttribute()
		method.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			method.class.constantPool)
	}
}

func (method *Method) calcArgSlotCount(paramTypes []string) {
	for _, paramType := range paramTypes {
		method.argSlotCount++
		if paramType == "J" || paramType == "D" {
			method.argSlotCount++
		}
	}
	if !method.IsStatic() {
		method.argSlotCount++
	}
}

func (method *Method) injectCodeAttribute(returnType string) {
	method.maxStack = 4
	method.maxLocals = method.argSlotCount
	switch returnType[0] {
	case 'V':
		method.code = []byte{0xfe, 0xb1} //return
	case 'L', '[':
		method.code = []byte{0xfe, 0xb0} //areturn
	case 'D':
		method.code = []byte{0xfe, 0xaf} //dreturn
	case 'F':
		method.code = []byte{0xfe, 0xae} //freturn
	case 'J':
		method.code = []byte{0xfe, 0xad} // lreturn
	default:
		method.code = []byte{0xfe, 0xac} // ireturn
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
func (method *Method) GetLineNumber(pc int) int {
	if method.IsNative() {
		return -2
	}
	if method.lineNumberTable == nil {
		return -1
	}
	return method.lineNumberTable.GetLineNumber(pc)
}

func (method *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := method.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlePc
	}
	return -1
}
