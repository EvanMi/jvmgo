package heap

import (
	"jvmgo/ch11/clazzfile"
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
	sourceFile        string
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool
	jClass            *Object
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
	class.sourceFile = getSourceFile(cf)
	return class
}

func getSourceFile(cf *clazzfile.ClassFile) string {
	if sfAttr := cf.SourceFileAttribute(); sfAttr != nil {
		return sfAttr.FileName()
	}
	return "Unknow"
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
func (class *Class) AccessFlags() uint16 {
	return class.accessFlags
}
func (class *Class) Name() string {
	return class.name
}
func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
}
func (class *Class) Fields() []*Field {
	return class.fields
}
func (class *Class) Methods() []*Method {
	return class.methods
}
func (class *Class) Loader() *ClassLoader {
	return class.loader
}
func (class *Class) SuperClass() *Class {
	return class.superClass
}
func (class *Class) InitStarted() bool {
	return class.initStarted
}
func (class *Class) JClass() *Object {
	return class.jClass
}
func (class *Class) StartInit() {
	class.initStarted = true
}
func (class *Class) StaticVars() Slots {
	return class.staticVars
}

func (class *Class) isAccessiableTo(other *Class) bool {
	return class.IsPublic() || class.GetPackageName() == other.GetPackageName()
}

func (class *Class) GetPackageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) GetMainMethod() *Method {
	return class.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

func (class *Class) GetClinitMethod() *Method {
	return class.GetStaticMethod("<clinit>", "()V")
}

func (class *Class) GetStaticMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, true)
}

func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic && method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := class; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

func (class *Class) isJlObject() bool {
	return class.name == "java/lang/Object"
}
func (class *Class) isJlCloneable() bool {
	return class.name == "java/lang/Cloneable"
}

func (class *Class) isJioSerializable() bool {
	return class.name == "java/io/Serializable"
}

func (class *Class) NewObject() *Object {
	return newObject(class)
}

func (class *Class) ArrayClass() *Class {
	arrayClassName := getArrayClassName(class.name)
	return class.loader.LoadClass(arrayClassName)
}
func (class *Class) JavaName() string {
	return strings.Replace(class.name, "/", ".", -1)
}
func (class *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[class.name]
	return ok
}

func (class *Class) GetInstanceMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, false)
}

func (class *Class) GetRefVar(fieldName, fieldDescriptor string) *Object {
	field := class.getField(fieldName, fieldDescriptor, true)
	return class.staticVars.GetRef(field.slotId)
}
func (class *Class) SetRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := class.getField(fieldName, fieldDescriptor, true)
	class.staticVars.SetRef(field.slotId, ref)
}
func (class *Class) SourceFile() string {
	return class.sourceFile
}
func (class *Class) GetConstructor(descriptor string) *Method {
	return class.GetInstanceMethod("<init>", descriptor)
}
func (class *Class) Interfaces() []*Class {
	return class.interfaces
}
func (class *Class) GetConstructors(publicOnly bool) []*Method {
	constructors := make([]*Method, 0, len(class.methods))
	for _, method := range class.methods {
		if method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				constructors = append(constructors, method)
			}
		}
	}
	return constructors
}
func (class *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(class.fields))
		for _, field := range class.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return class.fields
	}
}

func (class *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(class.methods))
	for _, method := range class.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}
