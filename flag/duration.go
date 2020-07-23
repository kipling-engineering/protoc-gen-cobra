package flag

import (
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/protobuf/types/known/durationpb"
)

type DurationValue struct {
	get func() *duration.Duration
	set func(*duration.Duration)
}

func NewDurationValue(get func() *duration.Duration, set func(*duration.Duration)) *DurationValue {
	return &DurationValue{get, set}
}

func (v *DurationValue) Set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	v.set(durationpb.New(d))
	return nil
}

func (v *DurationValue) Type() string { return "Duration" }

func (v *DurationValue) String() string {
	if w := v.get(); w != nil {
		return w.AsDuration().String()
	}
	return ""
}
