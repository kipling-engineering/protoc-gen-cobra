package flag

import (
	"fmt"
	"strconv"
	"strings"

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
	vals := any(T(0)).(protoreflect.Enum).Descriptor().Values()
	if i, err := strconv.ParseInt(val, 0, 32); err == nil {
		if n := vals.ByNumber(protoreflect.EnumNumber(i)); n != nil {
			return T(i), nil
		}
	} else if n := vals.ByName(protoreflect.Name(val)); n != nil {
		return T(int32(n.Number())), nil
	} else {
		for i := 0; i < vals.Len(); i++ {
			if n := vals.Get(i); strings.EqualFold(string(n.Name()), val) {
				return T(int32(n.Number())), nil
			}
		}
	}
	return 0, fmt.Errorf("unable to parse enum: %s", val)
}
