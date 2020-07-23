package flag

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type DoubleWrapperValue func(*wrappers.DoubleValue)

func NewDoubleWrapperValue(set func(*wrappers.DoubleValue)) *DoubleWrapperValue {
	v := DoubleWrapperValue(set)
	return &v
}

func (v *DoubleWrapperValue) Set(s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.Double(f))
	return nil
}

func (v *DoubleWrapperValue) Type() string { return "DoubleWrapper" }

func (v *DoubleWrapperValue) String() string { return "<nil>" }

type FloatWrapperValue func(*wrappers.FloatValue)

func NewFloatWrapperValue(set func(*wrappers.FloatValue)) *FloatWrapperValue {
	v := FloatWrapperValue(set)
	return &v
}

func (v *FloatWrapperValue) Set(s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.Float(float32(f)))
	return nil
}

func (v *FloatWrapperValue) Type() string { return "FloatWrapper" }

func (v *FloatWrapperValue) String() string { return "<nil>" }

type Int64WrapperValue func(*wrappers.Int64Value)

func NewInt64WrapperValue(set func(*wrappers.Int64Value)) *Int64WrapperValue {
	v := Int64WrapperValue(set)
	return &v
}

func (v *Int64WrapperValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.Int64(i))
	return nil
}

func (v *Int64WrapperValue) Type() string { return "Int64Wrapper" }

func (v *Int64WrapperValue) String() string { return "<nil>" }

type UInt64WrapperValue func(*wrappers.UInt64Value)

func NewUInt64WrapperValue(set func(*wrappers.UInt64Value)) *UInt64WrapperValue {
	v := UInt64WrapperValue(set)
	return &v
}

func (v *UInt64WrapperValue) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.UInt64(i))
	return nil
}

func (v *UInt64WrapperValue) Type() string { return "UInt64Wrapper" }

func (v *UInt64WrapperValue) String() string { return "<nil>" }

type Int32WrapperValue func(*wrappers.Int32Value)

func NewInt32WrapperValue(set func(*wrappers.Int32Value)) *Int32WrapperValue {
	v := Int32WrapperValue(set)
	return &v
}

func (v *Int32WrapperValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.Int32(int32(i)))
	return nil
}

func (v *Int32WrapperValue) Type() string { return "Int32Wrapper" }

func (v *Int32WrapperValue) String() string { return "<nil>" }

type UInt32WrapperValue func(*wrappers.UInt32Value)

func NewUInt32WrapperValue(set func(*wrappers.UInt32Value)) *UInt32WrapperValue {
	v := UInt32WrapperValue(set)
	return &v
}

func (v *UInt32WrapperValue) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.UInt32(uint32(i)))
	return nil
}

func (v *UInt32WrapperValue) Type() string { return "UInt32Wrapper" }

func (v *UInt32WrapperValue) String() string { return "<nil>" }

type BoolWrapperValue func(*wrappers.BoolValue)

func NewBoolWrapperValue(set func(*wrappers.BoolValue)) *BoolWrapperValue {
	v := BoolWrapperValue(set)
	return &v
}

func (v *BoolWrapperValue) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	(*v)(wrapperspb.Bool(b))
	return nil
}

func (v *BoolWrapperValue) Type() string { return "BoolWrapper" }

func (v *BoolWrapperValue) String() string { return "<nil>" }

type StringWrapperValue func(*wrappers.StringValue)

func NewStringWrapperValue(set func(*wrappers.StringValue)) *StringWrapperValue {
	v := StringWrapperValue(set)
	return &v
}

func (v *StringWrapperValue) Set(s string) error {
	(*v)(wrapperspb.String(s))
	return nil
}

func (v *StringWrapperValue) Type() string { return "StringWrapper" }

func (v *StringWrapperValue) String() string { return "<nil>" }

type BytesBase64WrapperValue func(*wrappers.BytesValue)

func NewBytesBase64WrapperValue(set func(*wrappers.BytesValue)) *BytesBase64WrapperValue {
	v := BytesBase64WrapperValue(set)
	return &v
}

func (v *BytesBase64WrapperValue) Set(s string) error {
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return err
	}
	(*v)(wrapperspb.Bytes(b))
	return nil
}

func (v *BytesBase64WrapperValue) Type() string { return "BytesBase64Wrapper" }

func (v *BytesBase64WrapperValue) String() string { return "<nil>" }
