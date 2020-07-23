package flag

import (
	"time"

	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/protobuf/types/known/durationpb"
)

type DurationValue func(*duration.Duration)

func NewDurationValue(set func(*duration.Duration)) *DurationValue {
	v := DurationValue(set)
	return &v
}

func (v *DurationValue) Set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	(*v)(durationpb.New(d))
	return nil
}

func (v *DurationValue) Type() string { return "Duration" }

func (v *DurationValue) String() string { return "<nil>" }
