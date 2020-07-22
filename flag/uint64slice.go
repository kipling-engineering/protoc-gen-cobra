package flag

import (
	"fmt"
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

func (s *uint64SliceValue) Type() string { return "uint64Slice" }

func (s *uint64SliceValue) String() string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = s.toString(d)
	}
	return "[" + strings.Join(out, ",") + "]"
}

func (s *uint64SliceValue) Append(val string) error {
	d, err := s.fromString(val)
	if err != nil {
		return err
	}
	*s.value = append(*s.value, d)
	return nil
}

func (s *uint64SliceValue) Replace(val []string) error {
	out := make([]uint64, len(val))
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

func (s *uint64SliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = s.toString(d)
	}
	return out
}

func (s *uint64SliceValue) fromString(val string) (uint64, error) {
	return strconv.ParseUint(val, 0, 64)
}

func (s *uint64SliceValue) toString(d uint64) string {
	return fmt.Sprintf("%d", d)
}
