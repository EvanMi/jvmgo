package heap

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (mdParser *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	mdParser.raw = descriptor
	mdParser.parsed = &MethodDescriptor{}
	mdParser.startParams()
	mdParser.parseParamTypes()
	mdParser.endParams()
	mdParser.parseReturnType()
	mdParser.finish()
	return mdParser.parsed
}

func (mdParser *MethodDescriptorParser) startParams() {
	if mdParser.readUint8() != '(' {
		mdParser.causePanic()
	}
}

func (mdParser *MethodDescriptorParser) endParams() {
	if mdParser.readUint8() != ')' {
		mdParser.causePanic()
	}
}

func (mdParser *MethodDescriptorParser) parseReturnType() {
	if mdParser.readUint8() == 'V' {
		mdParser.parsed.returnType = "V"
		return
	}
	mdParser.unreadUint8()
	t := mdParser.parseFieldType()
	if t != "" {
		mdParser.parsed.returnType = t
		return
	}
	mdParser.causePanic()
}

func (mdParser *MethodDescriptorParser) finish() {
	if mdParser.offset != len(mdParser.raw) {
		mdParser.causePanic()
	}
}

func (mdParser *MethodDescriptorParser) parseParamTypes() {
	for {
		t := mdParser.parseFieldType()
		if t != "" {
			mdParser.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (mdParser *MethodDescriptorParser) parseFieldType() string {
	switch mdParser.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return mdParser.parseObjectType()
	case '[':
		return mdParser.parseArrayType()
	default:
		mdParser.unreadUint8()
		return ""
	}
}

func (mdParser *MethodDescriptorParser) parseObjectType() string {
	unread := mdParser.raw[mdParser.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		mdParser.causePanic()
		return ""
	} else {
		objStart := mdParser.offset - 1
		objEnd := mdParser.offset + semicolonIndex + 1
		mdParser.offset = objEnd
		descriptor := mdParser.raw[objStart:objEnd]
		return descriptor
	}
}

func (mdParser *MethodDescriptorParser) parseArrayType() string {
	arrStart := mdParser.offset - 1
	mdParser.parseFieldType()
	arrEnd := mdParser.offset
	descriptor := mdParser.raw[arrStart:arrEnd]
	return descriptor
}

func (mdParser *MethodDescriptorParser) readUint8() uint8 {
	b := mdParser.raw[mdParser.offset]
	mdParser.offset++
	return b
}

func (mdParser *MethodDescriptorParser) unreadUint8() {
	mdParser.offset--
}

func (mdParser *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + mdParser.raw)
}
