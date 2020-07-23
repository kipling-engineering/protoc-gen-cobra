package flag

import (
	"fmt"
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
		var err error
		if out[i], err = s.fromString(v); err != nil {
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

func (s *uint32SliceValue) Type() string { return "uint32Slice" }

func (s *uint32SliceValue) String() string { return "[]" }

func (s *uint32SliceValue) Append(val string) error {
	d, err := s.fromString(val)
	if err != nil {
		return err
	}
	*s.value = append(*s.value, d)
	return nil
}

func (s *uint32SliceValue) Replace(val []string) error {
	out := make([]uint32, len(val))
	for i, v := range val {
		var err error
		out[i], err = s.fromString(v)
		if err != nil {
			return err
		}
	}
	*s.value = out
	return nil
}

func (s *uint32SliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = s.toString(d)
	}
	return out
}

func (s *uint32SliceValue) fromString(val string) (uint32, error) {
	d, err := strconv.ParseUint(val, 0, 32)
	if err != nil {
		return 0, err
	}
	return uint32(d), nil
}

func (s *uint32SliceValue) toString(d uint32) string {
	return fmt.Sprintf("%d", d)
}
