package heap

import "jvmgo/ch07/clazzfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func newFields(class *Class, cfFields []*clazzfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

func (field *Field) copyAttributes(cfField *clazzfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		field.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

func (field *Field) IsVolatile() bool {
	return field.accessFlags&ACC_VOLATILE != 0
}
func (field *Field) IsTransient() bool {
	return field.accessFlags&ACC_TRANSIENT != 0
}
func (field *Field) IsEnum() bool {
	return field.accessFlags&ACC_ENUM != 0
}

func (field *Field) ConstValueIndex() uint {
	return field.constValueIndex
}
func (field *Field) SlotId() uint {
	return field.slotId
}
func (field *Field) isLongOrDouble() bool {
	return field.descriptor == "J" || field.descriptor == "D"
}
