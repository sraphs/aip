package types

import (
	"google.golang.org/protobuf/reflect/protoreflect"
)

// NewEnumValueDescription produces an enum value description with the fully qualified enum value
// name and the enum value descriptor.
func NewEnumValueDescription(name string, desc protoreflect.EnumValueDescriptor) *EnumValueDescription {
	return &EnumValueDescription{
		enumValueName: name,
		desc:          desc,
	}
}

// EnumValueDescription maps a fully-qualified enum value name to its numeric value.
type EnumValueDescription struct {
	enumValueName string
	desc          protoreflect.EnumValueDescriptor
}

// Name returns the fully-qualified identifier name for the enum value.
func (ed *EnumValueDescription) Name() string {
	return ed.enumValueName
}

// Value returns the (numeric) value of the enum.
func (ed *EnumValueDescription) Value() int32 {
	return int32(ed.desc.Number())
}
