package flag

import (
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/durationpb"
)

type durationValue struct {
	set func(*duration.Duration)
}

func DurationVar(fs *pflag.FlagSet, p **duration.Duration, name, usage string) {
	fs.Var(&durationValue{func(d *duration.Duration) { *p = d }}, name, usage)
}

func (v *durationValue) Set(val string) error {
	if d, err := parseDuration(val); err != nil {
		return err
	} else {
		v.set(d)
		return nil
	}
}

func (*durationValue) Type() string { return "duration" }

func (*durationValue) String() string { return "<nil>" }

type durationSliceValue struct {
	set func([]*duration.Duration)
}

func DurationSliceVar(fs *pflag.FlagSet, p *[]*duration.Duration, name, usage string) {
	var changed bool
	set := func(out []*duration.Duration) {
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	fs.Var(&durationSliceValue{set}, name, usage)
}

func (s *durationSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]*duration.Duration, len(ss))
	for i, v := range ss {
		var err error
		if out[i], err = parseDuration(v); err != nil {
			return err
		}
	}
	s.set(out)
	return nil
}

func (s *durationSliceValue) Type() string { return "durationSlice" }

func (s *durationSliceValue) String() string { return "[]" }

func parseDuration(val string) (*duration.Duration, error) {
	if d, err := time.ParseDuration(val); err != nil {
		return nil, err
	} else {
		return durationpb.New(d), nil
	}
}
