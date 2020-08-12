package flag

import (
	"strings"
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

type timestampSliceValue struct {
	set func([]*timestamp.Timestamp)
}

func TimestampSliceVar(fs *pflag.FlagSet, p *[]*timestamp.Timestamp, name, usage string) {
	var changed bool
	set := func(out []*timestamp.Timestamp) {
		if !changed {
			*p = out
			changed = true
		} else {
			*p = append(*p, out...)
		}
	}
	fs.Var(&timestampSliceValue{set}, name, usage)
}

func (s *timestampSliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([]*timestamp.Timestamp, len(ss))
	for i, v := range ss {
		var err error
		if out[i], err = parseTimestamp(v); err != nil {
			return err
		}
	}
	s.set(out)
	return nil
}

func (s *timestampSliceValue) Type() string { return "timestampSlice" }

func (s *timestampSliceValue) String() string { return "[]" }

func ParseTimestamp(val string) (interface{}, error) { return parseTimestamp(val) }

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
