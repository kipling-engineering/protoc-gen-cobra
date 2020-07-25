package flag

import (
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

type uint64SliceValue struct {
	value   *[]uint64
	changed bool
}

func Uint64SliceVar(fs *pflag.FlagSet, p *[]uint64, name, usage string) {
	fs.Var(&uint64SliceValue{value: p}, name, usage)
}

func (s *uint64SliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]uint64, len(ss))
	for i, v := range ss {
		var err error
		if out[i], err = strconv.ParseUint(v, 0, 64); err != nil {
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

func (*uint64SliceValue) Type() string { return "uint64Slice" }

func (*uint64SliceValue) String() string { return "[]" }
