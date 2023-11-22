package heap

import (
	"jvmgo/ch06/clazzfile"
	"strings"
)

type Class struct {
	accessFlags       uint16
	name              string
	superClassName    string
	interfaceNames    []string
	constantPool      *ConstantPool
	fields            []*Field
	methods           []*Method
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func newClass(cf *clazzfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (class *Class) IsPublic() bool {
	return class.accessFlags&ACC_PUBLIC != 0
}
func (class *Class) IsFinal() bool {
	return class.accessFlags&ACC_FINAL != 0
}
func (class *Class) IsSuper() bool {
	return class.accessFlags&ACC_SUPER != 0
}
func (class *Class) IsInterface() bool {
	return class.accessFlags&ACC_INTERFACE != 0
}
func (class *Class) IsAbstract() bool {
	return class.accessFlags&ACC_ABSTRACT != 0
}
func (class *Class) IsSynthetic() bool {
	return class.accessFlags&ACC_SYNTHETIC != 0
}
func (class *Class) IsAnnotation() bool {
	return class.accessFlags&ACC_ANNOTATION != 0
}
func (class *Class) IsEnum() bool {
	return class.accessFlags&ACC_ENUM != 0
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) StaticVars() Slots {
	return class.staticVars
}

func (class *Class) isAccessiableTo(other *Class) bool {
	return class.IsPublic() || class.getPackageName() == other.getPackageName()
}

func (class *Class) getPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) GetMainMethod() *Method {
	return class.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}
