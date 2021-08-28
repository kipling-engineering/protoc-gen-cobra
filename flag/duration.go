package flag

import (
	"strings"

	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/NathanBaulch/protoc-gen-cobra/ptypes"
)

func DurationVar(fs *pflag.FlagSet, p **durationpb.Duration, name, usage string) {
	v := fs.String(name, "", usage)
	WithPostSetHookE(fs, name, func() (err error) { *p, err = ptypes.ToDuration(v); return })
}

type durationSliceValue struct {
	value   *[]*durationpb.Duration
	changed bool
}

func DurationSliceVar(fs *pflag.FlagSet, p *[]*durationpb.Duration, name, usage string) {
	fs.Var(&durationSliceValue{value: p}, name, usage)
}

func (s *durationSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]*durationpb.Duration, len(ss))
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
