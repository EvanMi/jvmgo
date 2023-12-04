package heap

func (int8Arr *Object) Bytes() []int8 {
	return int8Arr.data.([]int8)
}

func (int16Arr *Object) Shorts() []int16 {
	return int16Arr.data.([]int16)
}

func (intArr *Object) Ints() []int32 {
	return intArr.data.([]int32)
}

func (longArr *Object) Longs() []int64 {
	return longArr.data.([]int64)
}

func (charArr *Object) Chars() []uint16 {
	return charArr.data.([]uint16)
}

func (floatArr *Object) Floats() []float32 {
	return floatArr.data.([]float32)
}
func (doubleArr *Object) Doubles() []float64 {
	return doubleArr.data.([]float64)
}
func (refArr *Object) Refs() []*Object {
	return refArr.data.([]*Object)
}

func (arr *Object) ArrayLength() int32 {
	switch arr.data.(type) {
	case []int8:
		return int32(len(arr.data.([]int8)))
	case []int16:
		return int32(len(arr.data.([]int16)))
	case []int32:
		return int32(len(arr.data.([]int32)))
	case []int64:
		return int32(len(arr.data.([]int64)))
	case []uint16:
		return int32(len(arr.data.([]uint16)))
	case []float32:
		return int32(len(arr.data.([]float32)))
	case []float64:
		return int32(len(arr.data.([]float64)))
	case []*Object:
		return int32(len(arr.data.([]*Object)))
	default:
		panic("Not array!")
	}
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dst.data.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dst.data.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos+length]
		_dst := dst.data.([]int64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dst.data.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dst.data.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}
