package flag

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type DoubleWrapperValue struct {
	get func() *wrappers.DoubleValue
	set func(*wrappers.DoubleValue)
}

func NewDoubleWrapperValue(get func() *wrappers.DoubleValue, set func(*wrappers.DoubleValue)) *DoubleWrapperValue {
	return &DoubleWrapperValue{get, set}
}

func (v *DoubleWrapperValue) Set(s string) error {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	v.set(wrapperspb.Double(f))
	return nil
}

func (v *DoubleWrapperValue) Type() string { return "DoubleWrapper" }

func (v *DoubleWrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatFloat(w.Value, 'g', -1, 64)
	}
	return ""
}

type FloatWrapperValue struct {
	get func() *wrappers.FloatValue
	set func(*wrappers.FloatValue)
}

func NewFloatWrapperValue(get func() *wrappers.FloatValue, set func(*wrappers.FloatValue)) *FloatWrapperValue {
	return &FloatWrapperValue{get, set}
}

func (v *FloatWrapperValue) Set(s string) error {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	v.set(wrapperspb.Float(float32(f)))
	return nil
}

func (v *FloatWrapperValue) Type() string { return "FloatWrapper" }

func (v *FloatWrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatFloat(float64(w.Value), 'g', -1, 32)
	}
	return ""
}

type Int64WrapperValue struct {
	get func() *wrappers.Int64Value
	set func(*wrappers.Int64Value)
}

func NewInt64WrapperValue(get func() *wrappers.Int64Value, set func(*wrappers.Int64Value)) *Int64WrapperValue {
	return &Int64WrapperValue{get, set}
}

func (v *Int64WrapperValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	v.set(wrapperspb.Int64(i))
	return nil
}

func (v *Int64WrapperValue) Type() string { return "Int64Wrapper" }

func (v *Int64WrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatInt(w.Value, 10)
	}
	return ""
}

type UInt64WrapperValue struct {
	get func() *wrappers.UInt64Value
	set func(*wrappers.UInt64Value)
}

func NewUInt64WrapperValue(get func() *wrappers.UInt64Value, set func(*wrappers.UInt64Value)) *UInt64WrapperValue {
	return &UInt64WrapperValue{get, set}
}

func (v *UInt64WrapperValue) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return err
	}
	v.set(wrapperspb.UInt64(i))
	return nil
}

func (v *UInt64WrapperValue) Type() string { return "UInt64Wrapper" }

func (v *UInt64WrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatUint(w.Value, 10)
	}
	return ""
}

type Int32WrapperValue struct {
	get func() *wrappers.Int32Value
	set func(*wrappers.Int32Value)
}

func NewInt32WrapperValue(get func() *wrappers.Int32Value, set func(*wrappers.Int32Value)) *Int32WrapperValue {
	return &Int32WrapperValue{get, set}
}

func (v *Int32WrapperValue) Set(s string) error {
	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return err
	}
	v.set(wrapperspb.Int32(int32(i)))
	return nil
}

func (v *Int32WrapperValue) Type() string { return "Int32Wrapper" }

func (v *Int32WrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatInt(int64(w.Value), 10)
	}
	return ""
}

type UInt32WrapperValue struct {
	get func() *wrappers.UInt32Value
	set func(*wrappers.UInt32Value)
}

func NewUInt32WrapperValue(get func() *wrappers.UInt32Value, set func(*wrappers.UInt32Value)) *UInt32WrapperValue {
	return &UInt32WrapperValue{get, set}
}

func (v *UInt32WrapperValue) Set(s string) error {
	i, err := strconv.ParseUint(s, 0, 32)
	if err != nil {
		return err
	}
	v.set(wrapperspb.UInt32(uint32(i)))
	return nil
}

func (v *UInt32WrapperValue) Type() string { return "UInt32Wrapper" }

func (v *UInt32WrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatUint(uint64(w.Value), 10)
	}
	return ""
}

type BoolWrapperValue struct {
	get func() *wrappers.BoolValue
	set func(*wrappers.BoolValue)
}

func NewBoolWrapperValue(get func() *wrappers.BoolValue, set func(*wrappers.BoolValue)) *BoolWrapperValue {
	return &BoolWrapperValue{get, set}
}

func (v *BoolWrapperValue) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	v.set(wrapperspb.Bool(b))
	return nil
}

func (v *BoolWrapperValue) Type() string { return "BoolWrapper" }

func (v *BoolWrapperValue) String() string {
	if w := v.get(); w != nil {
		return strconv.FormatBool(w.Value)
	}
	return ""
}

type StringWrapperValue struct {
	get func() *wrappers.StringValue
	set func(*wrappers.StringValue)
}

func NewStringWrapperValue(get func() *wrappers.StringValue, set func(*wrappers.StringValue)) *StringWrapperValue {
	return &StringWrapperValue{get, set}
}

func (v *StringWrapperValue) Set(s string) error {
	v.set(wrapperspb.String(s))
	return nil
}

func (v *StringWrapperValue) Type() string { return "StringWrapper" }

func (v *StringWrapperValue) String() string {
	if w := v.get(); w != nil {
		return w.Value
	}
	return ""
}

type BytesBase64WrapperValue struct {
	get func() *wrappers.BytesValue
	set func(*wrappers.BytesValue)
}

func NewBytesBase64WrapperValue(get func() *wrappers.BytesValue, set func(*wrappers.BytesValue)) *BytesBase64WrapperValue {
	return &BytesBase64WrapperValue{get, set}
}

func (v *BytesBase64WrapperValue) Set(s string) error {
	b, err := base64.StdEncoding.DecodeString(strings.TrimSpace(s))
	if err != nil {
		return err
	}
	v.set(wrapperspb.Bytes(b))
	return nil
}

func (v *BytesBase64WrapperValue) Type() string { return "BytesBase64Wrapper" }

func (v *BytesBase64WrapperValue) String() string {
	if w := v.get(); w != nil {
		return base64.StdEncoding.EncodeToString(w.Value)
	}
	return ""
}
