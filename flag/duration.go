package flag

import (
	"strings"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/spf13/pflag"

	"github.com/NathanBaulch/protoc-gen-cobra/ptypes"
)

func DurationVar(fs *pflag.FlagSet, p **duration.Duration, name, usage string) {
	v := fs.String(name, "", usage)
	hook := func() error {
		if d, err := ptypes.ToDuration(v); err != nil {
			return err
		} else {
			*p = d
			return nil
		}
	}
	WithPostSetHookE(fs, name, hook)
}

type durationSliceValue struct {
	value   *[]*duration.Duration
	changed bool
}

func DurationSliceVar(fs *pflag.FlagSet, p *[]*duration.Duration, name, usage string) {
	fs.Var(&durationSliceValue{value: p}, name, usage)
}

func (s *durationSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]*duration.Duration, len(ss))
	for i, v := range ss {
		var err error
		if out[i], err = ptypes.ToDuration(v); err != nil {
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

func (*durationSliceValue) Type() string { return "durationSlice" }

func (*durationSliceValue) String() string { return "[]" }

func ParseDuration(val string) (interface{}, error) { return ptypes.ToDuration(val) }
