package heap

func (class *Class) isAssignableFrom(other *Class) bool {
	s, t := other, class

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.isSubClassOf(t)
	} else {
		return s.isImplements(t)
	}
}

func (class *Class) isSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (class *Class) isImplements(iface *Class) bool {
	for c := class; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range class.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}
