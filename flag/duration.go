package flag

import (
	"time"
	"unsafe"

	"github.com/golang/protobuf/ptypes/duration"
	"google.golang.org/protobuf/types/known/durationpb"
)

type DurationValue struct {
	ptr unsafe.Pointer
}

func NewDurationValue(value **duration.Duration) *DurationValue {
	return &DurationValue{unsafe.Pointer(value)}
}

func (v *DurationValue) Set(s string) error {
	d, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	*(**duration.Duration)(v.ptr) = durationpb.New(d)
	return nil
}

func (v *DurationValue) Type() string { return "Duration" }

func (v *DurationValue) String() string { return "<nil>" }
