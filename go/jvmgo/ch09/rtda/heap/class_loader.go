package heap

import (
	"fmt"
	"jvmgo/ch09/clazzfile"
	"jvmgo/ch09/clazzpath"
)

type ClassLoader struct {
	cp          *clazzpath.Classpath
	verboseFlag bool
	classMap    map[string]*Class
}

func NewClassLoader(cp *clazzpath.Classpath, verboseFlag bool) *ClassLoader {
	loader := &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
	loader.loadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

func (classLoader *ClassLoader) loadBasicClasses() {
	jlClassClass := classLoader.LoadClass("java/lang/Class")
	for _, class := range classLoader.classMap {
		if class.jClass == nil {
			class.jClass = jlClassClass.NewObject()
			class.jClass.extra = class
		}
	}
}

func (classLoader *ClassLoader) loadPrimitiveClasses() {
	for primitiveType := range primitiveTypes {
		classLoader.loadPrimitiveClass(primitiveType)
	}
}

func (classLoader *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        className,
		loader:      classLoader,
		initStarted: true,
	}
	class.jClass = classLoader.classMap["java/lang/Class"].NewObject()
	class.jClass.extra = class
	classLoader.classMap[className] = class
}

func (classLoader *ClassLoader) LoadClass(name string) *Class {
	if class, ok := classLoader.classMap[name]; ok {
		return class
	}
	var class *Class
	if name[0] == '[' {
		// array class
		class = classLoader.loadArrayClass(name)
	} else {
		class = classLoader.loadNonArrayClass(name)
	}

	if jlClassClass, ok := classLoader.classMap["java/lang/Class"]; ok {
		class.jClass = jlClassClass.NewObject()
		class.jClass.extra = class
	}
	return class
}

func (classLoader *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name:        name,
		loader:      classLoader,
		initStarted: true,
		superClass:  classLoader.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			classLoader.LoadClass("java/lang/Cloneable"),
			classLoader.LoadClass("java/io/Serializable"),
		},
	}
	classLoader.classMap[name] = class
	return class
}

func (classLoader *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := classLoader.readClass(name)
	class := classLoader.defineClass(data)
	link(class)
	if classLoader.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

func (classLoader *ClassLoader) readClass(name string) ([]byte, clazzpath.Entry) {
	data, entry, err := classLoader.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name + "--" + err.Error())
	}
	return data, entry
}

func (classLoader *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = classLoader
	resolveSuperClass(class)
	resolveInterfaces(class)
	classLoader.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := clazzfile.Parse(data)
	if err != nil {
		panic(err)
	}
	return newClass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	//pass
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class) {
	slodId := uint(0)
	if class.superClass != nil {
		slodId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slodId
			slodId++
			if field.isLongOrDouble() {
				slodId++
			}
		}
	}
	class.instanceSlotCount = slodId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.slotId

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}

	}
}
