package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (methodDescriptor *MethodDescriptor) addParameterType(t string) {
	pLen := len(methodDescriptor.parameterTypes)
	if pLen == cap(methodDescriptor.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, methodDescriptor.parameterTypes)
		methodDescriptor.parameterTypes = s
	}
	methodDescriptor.parameterTypes = append(methodDescriptor.parameterTypes, t)
}
