package clazzfile

import "fmt"

type ClassFile struct {
	//magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (classFile *ClassFile) read(reader *ClassReader) {
	classFile.readAndCheckMagic(reader)
	classFile.readAndCheckVersion(reader)
	classFile.constantPool = readContantPool(reader)
	classFile.accessFlags = reader.readUint16()
	classFile.thisClass = reader.readUint16()
	classFile.superClass = reader.readUint16()
	classFile.interfaces = reader.readUint16s()
	classFile.fields = readMembers(reader, classFile.constantPool)
	classFile.methods = readMembers(reader, classFile.constantPool)
	classFile.attributes = readAttributes(reader, classFile.constantPool)
}

func (classFile *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (classFile *ClassFile) readAndCheckVersion(reader *ClassReader) {
	classFile.minorVersion = reader.readUint16()
	classFile.majorVersion = reader.readUint16()
	switch classFile.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if classFile.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func (classFile *ClassFile) MinorVersion() uint16 {
	return classFile.minorVersion
}
func (classFile *ClassFile) MajorVersion() uint16 {
	return classFile.majorVersion
}
func (classFile *ClassFile) ConstantPool() ConstantPool {
	return classFile.constantPool
}
func (classFile *ClassFile) AccessFlags() uint16 {
	return classFile.accessFlags
}
func (classFile *ClassFile) Fields() []*MemberInfo {
	return classFile.fields
}
func (classFile *ClassFile) Methods() []*MemberInfo {
	return classFile.methods
}

func (classFile *ClassFile) ClassName() string {
	return classFile.constantPool.getClassName(classFile.thisClass)
}

func (classFile *ClassFile) SuperClassName() string {
	if classFile.superClass > 0 {
		return classFile.constantPool.getClassName(classFile.superClass)
	}
	return ""
}

func (classFile *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(classFile.interfaces))
	for i, cpIndex := range classFile.interfaces {
		interfaceNames[i] = classFile.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
