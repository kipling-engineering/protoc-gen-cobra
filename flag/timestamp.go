package flag

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/pflag"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type timestampValue struct {
	set func(*timestamp.Timestamp)
}

func TimestampVar(fs *pflag.FlagSet, p **timestamp.Timestamp, name, usage string) {
	fs.Var(&timestampValue{func(d *timestamp.Timestamp) { *p = d }}, name, usage)
}

func (v *timestampValue) Set(val string) error {
	if d, err := parseTimestamp(val); err != nil {
		return err
	} else {
		v.set(d)
		return nil
	}
}

func (*timestampValue) Type() string { return "timestamp" }

func (*timestampValue) String() string { return "<nil>" }

func parseTimestamp(val string) (*timestamp.Timestamp, error) {
	var err error
	for _, layout := range []string{
		"2006-01-02T15:04:05.999999999Z07:00",
		"2006-01-02T15:04:05.999999999Z07",
		"2006-01-02T15:04:05.999999999",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z07",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02T15",
		"2006-01-02",
	} {
		var t time.Time
		t, err = time.Parse(layout, val)
		if err != nil {
			continue
		}
		return timestamppb.New(t), nil
	}
	return nil, err
}
