package heap

func (class *Class) IsArray() bool {
	return class.name[0] == '['
}

func (class *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(class.name)
	return class.loader.LoadClass(componentClassName)
}

func (class *Class) NewArray(count uint) *Object {
	if !class.IsArray() {
		panic("Not array class: " + class.name)
	}
	switch class.Name() {
	case "[Z":
		return &Object{class, make([]int8, count)}
	case "[B":
		return &Object{class, make([]int8, count)}
	case "[C":
		return &Object{class, make([]uint16, count)}
	case "[S":
		return &Object{class, make([]int16, count)}
	case "[I":
		return &Object{class, make([]int32, count)}
	case "[J":
		return &Object{class, make([]int64, count)}
	case "[F":
		return &Object{class, make([]float32, count)}
	case "[D":
		return &Object{class, make([]float64, count)}
	default:
		return &Object{class, make([]*Object, count)}
	}
}
