package heap

import (
	"fmt"
	"jvmgo/ch09/clazzfile"
)

type Constant interface{}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func (pool *ConstantPool) GetConstant(index uint) Constant {
	if c := pool.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("no constants at index %d", index))
}

func newConstantPool(class *Class, cfCp clazzfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{
		class:  class,
		consts: consts,
	}
	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo := cpInfo.(type) {
		case *clazzfile.ConstantIntegerInfo:
			consts[i] = cpInfo.Value()
		case *clazzfile.ConstantFloatInfo:
			consts[i] = cpInfo.Value()
		case *clazzfile.ConstantLongInfo:
			consts[i] = cpInfo.Value()
			i++
		case *clazzfile.ConstantDoubleInfo:
			consts[i] = cpInfo.Value()
			i++
		case *clazzfile.ConstantStringInfo:
			consts[i] = cpInfo.String()
		case *clazzfile.ConstantClassInfo:
			consts[i] = newClassRef(rtCp, cpInfo)
		case *clazzfile.ConstantFieldrefInfo:
			consts[i] = newFieldRef(rtCp, cpInfo)
		case *clazzfile.ConstantMethodrefInfo:
			consts[i] = newMethodRef(rtCp, cpInfo)
		case *clazzfile.ConstantInterfaceMethodrefInfo:
			consts[i] = newInterfaceMethodRef(rtCp, cpInfo)
		default:
			// todo
		}
	}
	return rtCp
}
