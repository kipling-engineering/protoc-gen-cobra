package flag

import (
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

func parseDuration(val string) (*duration.Duration, error) {
	if d, err := time.ParseDuration(val); err != nil {
		return nil, err
	} else {
		return durationpb.New(d), nil
	}
}
