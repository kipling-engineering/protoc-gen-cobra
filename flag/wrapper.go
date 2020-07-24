package flag

import (
	"encoding/base64"
	"strconv"
	"strings"
	"unsafe"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type DoubleWrapperValue struct {
	ptr unsafe.Pointer
}

func NewDoubleWrapperValue(value **wrappers.DoubleValue) *DoubleWrapperValue {
	return &DoubleWrapperValue{unsafe.Pointer(value)}
}

func (v *DoubleWrapperValue) Set(s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*(**wrappers.DoubleValue)(v.ptr) = wrapperspb.Double(f)
	return nil
}

func (v *DoubleWrapperValue) Type() string { return "DoubleWrapper" }

func (v *DoubleWrapperValue) String() string { return "<nil>" }

type FloatWrapperValue struct {
	ptr unsafe.Pointer
}

func NewFloatWrapperValue(value **wrappers.FloatValue) *FloatWrapperValue {
	return &FloatWrapperValue{unsafe.Pointer(value)}
}

func (v *FloatWrapperValue) Set(s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*(**wrappers.FloatValue)(v.ptr) = wrapperspb.Float(float32(f))
	return nil
}

func (v *FloatWrapperValue) Type() string { return "FloatWrapper" }

func (v *FloatWrapperValue) String() string { return "<nil>" }

type Int64WrapperValue struct {
	ptr unsafe.Pointer
}

func NewInt64WrapperValue(value **wrappers.Int64Value) *Int64WrapperValue {
	return &Int64WrapperValue{unsafe.Pointer(value)}
}

func (v *Int64WrapperValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	*(**wrappers.Int64Value)(v.ptr) = wrapperspb.Int64(i)
	return nil
}

func (v *Int64WrapperValue) Type() string { return "Int64Wrapper" }

func (v *Int64WrapperValue) String() string { return "<nil>" }

type UInt64WrapperValue struct {
	ptr unsafe.Pointer
}

func NewUInt64WrapperValue(value **wrappers.UInt64Value) *UInt64WrapperValue {
	return &UInt64WrapperValue{unsafe.Pointer(value)}
}

func (v *UInt64WrapperValue) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return err
	}
	*(**wrappers.UInt64Value)(v.ptr) = wrapperspb.UInt64(i)
	return nil
}

func (v *UInt64WrapperValue) Type() string { return "UInt64Wrapper" }

func (v *UInt64WrapperValue) String() string { return "<nil>" }

type Int32WrapperValue struct {
	ptr unsafe.Pointer
}

func NewInt32WrapperValue(value **wrappers.Int32Value) *Int32WrapperValue {
	return &Int32WrapperValue{unsafe.Pointer(value)}
}

func (v *Int32WrapperValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return err
	}
	*(**wrappers.Int32Value)(v.ptr) = wrapperspb.Int32(int32(i))
	return nil
}

func (v *Int32WrapperValue) Type() string { return "Int32Wrapper" }

func (v *Int32WrapperValue) String() string { return "<nil>" }

type UInt32WrapperValue struct {
	ptr unsafe.Pointer
}

func NewUInt32WrapperValue(value **wrappers.UInt32Value) *UInt32WrapperValue {
	return &UInt32WrapperValue{unsafe.Pointer(value)}
}

func (v *UInt32WrapperValue) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return err
	}
	*(**wrappers.UInt32Value)(v.ptr) = wrapperspb.UInt32(uint32(i))
	return nil
}

func (v *UInt32WrapperValue) Type() string { return "UInt32Wrapper" }

func (v *UInt32WrapperValue) String() string { return "<nil>" }

type BoolWrapperValue struct {
	ptr unsafe.Pointer
}

func NewBoolWrapperValue(value **wrappers.BoolValue) *BoolWrapperValue {
	return &BoolWrapperValue{unsafe.Pointer(value)}
}

func (v *BoolWrapperValue) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*(**wrappers.BoolValue)(v.ptr) = wrapperspb.Bool(b)
	return nil
}

func (v *BoolWrapperValue) Type() string { return "BoolWrapper" }

func (v *BoolWrapperValue) String() string { return "<nil>" }

type StringWrapperValue struct {
	ptr unsafe.Pointer
}

func NewStringWrapperValue(value **wrappers.StringValue) *StringWrapperValue {
	return &StringWrapperValue{unsafe.Pointer(value)}
}

func (v *StringWrapperValue) Set(s string) error {
	*(**wrappers.StringValue)(v.ptr) = wrapperspb.String(s)
	return nil
}

func (v *StringWrapperValue) Type() string { return "StringWrapper" }

func (v *StringWrapperValue) String() string { return "<nil>" }

type BytesBase64WrapperValue struct {
	ptr unsafe.Pointer
}

func NewBytesBase64WrapperValue(value **wrappers.BytesValue) *BytesBase64WrapperValue {
	return &BytesBase64WrapperValue{unsafe.Pointer(value)}
}

func (v *BytesBase64WrapperValue) Set(s string) error {
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return err
	}
	*(**wrappers.BytesValue)(v.ptr) = wrapperspb.Bytes(b)
	return nil
}

func (v *BytesBase64WrapperValue) Type() string { return "BytesBase64Wrapper" }

func (v *BytesBase64WrapperValue) String() string { return "<nil>" }
