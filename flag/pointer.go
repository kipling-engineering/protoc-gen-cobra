package flag

import "github.com/spf13/pflag"

func BoolPointerVar(fs *pflag.FlagSet, p **bool, name, usage string) {
	v := fs.Bool(name, false, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func Int32PointerVar(fs *pflag.FlagSet, p **int32, name, usage string) {
	v := fs.Int32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func Int64PointerVar(fs *pflag.FlagSet, p **int64, name, usage string) {
	v := fs.Int64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func Uint32PointerVar(fs *pflag.FlagSet, p **uint32, name, usage string) {
	v := fs.Uint32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func Uint64PointerVar(fs *pflag.FlagSet, p **uint64, name, usage string) {
	v := fs.Uint64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func Float32PointerVar(fs *pflag.FlagSet, p **float32, name, usage string) {
	v := fs.Float32(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func Float64PointerVar(fs *pflag.FlagSet, p **float64, name, usage string) {
	v := fs.Float64(name, 0, usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

func StringPointerVar(fs *pflag.FlagSet, p **string, name, usage string) {
	v := fs.String(name, "", usage)
	f := fs.Lookup(name)
	f.Value = &pointerValue{f.Value, func() { *p = v }}
}

type pointerValue struct {
	pflag.Value
	set func()
}

func (v *pointerValue) Set(s string) error {
	if err := v.Value.Set(s); err != nil {
		return err
	}
	v.set()
	return nil
}

func (v *pointerValue) Type() string { return v.Value.Type() + "Pointer" }

func (*pointerValue) String() string { return "<nil>" }
