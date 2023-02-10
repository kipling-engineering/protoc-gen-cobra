package flag

import (
	"strings"

	"github.com/spf13/pflag"
)

type sliceValue[T any] struct {
	value   *[]T
	changed bool
	parser  func(val string) (T, error)
}

func SliceVar[T any](fs *pflag.FlagSet, parser func(val string) (T, error), p *[]T, name, usage string) {
	fs.Var(&sliceValue[T]{value: p, parser: parser}, name, usage)
}

func (s *sliceValue[T]) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]T, len(ss))
	for i, v := range ss {
		var err error
		if out[i], err = s.parser(v); err != nil {
			return err
		}
	}
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (*sliceValue[T]) Type() string { return "slice" }

func (*sliceValue[T]) String() string { return "[]" }

func Uint32SliceVar(fs *pflag.FlagSet, p *[]uint32, name, usage string) {
	SliceVar[uint32](fs, ParseUint32E, p, name, usage)
}

func Uint64SliceVar(fs *pflag.FlagSet, p *[]uint64, name, usage string) {
	SliceVar[uint64](fs, ParseUint64E, p, name, usage)
}
