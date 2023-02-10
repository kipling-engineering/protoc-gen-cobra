package flag

import (
	"strconv"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func EnumVar[T ~int32](fs *pflag.FlagSet, p *T, name, usage string) {
	v := fs.String(name, "", usage)
	WithPostSetHookE(fs, name, func() (err error) { *p, err = ParseEnumE[T](*v); return })
}

func EnumPointerVar[T ~int32](fs *pflag.FlagSet, p **T, name, usage string) {
	v := fs.String(name, "", usage)
	WithPostSetHookE(fs, name, func() error {
		if e, err := ParseEnumE[T](*v); err != nil {
			return err
		} else {
			*p = &e
			return nil
		}
	})
}

func EnumSliceVar[T ~int32](fs *pflag.FlagSet, p *[]T, name, usage string) {
	SliceVar[T](fs, ParseEnumE[T], p, name, usage)
}

func ParseEnumE[T ~int32](val string) (T, error) {
	if n := any(T(0)).(protoreflect.Enum).Descriptor().Values().ByName(protoreflect.Name(val)); n != nil {
		return T(int32(n.Number())), nil
	} else if i, err := strconv.ParseInt(val, 0, 32); err == nil {
		return T(i), nil
	} else {
		return 0, err
	}
}
