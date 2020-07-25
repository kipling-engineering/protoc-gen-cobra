package flag

import (
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type uint32SliceValue struct {
	value   *[]uint32
	changed bool
}

func Uint32SliceVar(fs *pflag.FlagSet, p *[]uint32, name, usage string) {
	fs.Var(&uint32SliceValue{value: p}, name, usage)
}

func (s *uint32SliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]uint32, len(ss))
	for i, v := range ss {
		d, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return err
		}
		out[i] = uint32(d)
	}
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (*uint32SliceValue) Type() string { return "uint32Slice" }

func (*uint32SliceValue) String() string { return "[]" }
