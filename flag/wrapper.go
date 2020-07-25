package flag

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func BoolWrapperVar(fs *pflag.FlagSet, p **wrappers.BoolValue, name, usage string) {
	v := fs.Bool(name, false, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Bool(*v) }}
}

func Int32WrapperVar(fs *pflag.FlagSet, p **wrappers.Int32Value, name, usage string) {
	v := fs.Int32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Int32(*v) }}
}

func Int64WrapperVar(fs *pflag.FlagSet, p **wrappers.Int64Value, name, usage string) {
	v := fs.Int64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Int64(*v) }}
}

func UInt32WrapperVar(fs *pflag.FlagSet, p **wrappers.UInt32Value, name, usage string) {
	v := fs.Uint32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.UInt32(*v) }}
}

func UInt64WrapperVar(fs *pflag.FlagSet, p **wrappers.UInt64Value, name, usage string) {
	v := fs.Uint64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.UInt64(*v) }}
}

func FloatWrapperVar(fs *pflag.FlagSet, p **wrappers.FloatValue, name, usage string) {
	v := fs.Float32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Float(*v) }}
}

func DoubleWrapperVar(fs *pflag.FlagSet, p **wrappers.DoubleValue, name, usage string) {
	v := fs.Float64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Double(*v) }}
}

func StringWrapperVar(fs *pflag.FlagSet, p **wrappers.StringValue, name, usage string) {
	v := fs.String(name, "", usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.String(*v) }}
}

func BytesBase64WrapperVar(fs *pflag.FlagSet, p **wrappers.BytesValue, name, usage string) {
	v := fs.BytesBase64(name, nil, usage)
	f := fs.Lookup(name)
	f.Value = &wrapperValue{f.Value, func() { *p = wrapperspb.Bytes(*v) }}
}

type wrapperValue struct {
	pflag.Value
	set func()
}

func (v *wrapperValue) Set(s string) error {
	if err := v.Value.Set(s); err != nil {
		return err
	}
	v.set()
	return nil
}

func (v *wrapperValue) Type() string { return v.Value.Type() + "Wrapper" }

func (*wrapperValue) String() string { return "<nil>" }
