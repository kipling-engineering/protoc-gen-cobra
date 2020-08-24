package flag

import (
	"strings"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/pflag"

	"github.com/NathanBaulch/protoc-gen-cobra/ptypes"
)

type timestampValue struct {
	set func(*timestamp.Timestamp)
}

func TimestampVar(fs *pflag.FlagSet, p **timestamp.Timestamp, name, usage string) {
	fs.Var(&timestampValue{func(d *timestamp.Timestamp) { *p = d }}, name, usage)
}

func (v *timestampValue) Set(val string) error {
	if d, err := ptypes.ToTimestamp(val); err != nil {
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
		if out[i], err = ptypes.ToTimestamp(v); err != nil {
			return err
		}
	}
	s.set(out)
	return nil
}

func (s *timestampSliceValue) Type() string { return "timestampSlice" }

func (s *timestampSliceValue) String() string { return "[]" }

func ParseTimestamp(val string) (interface{}, error) { return ptypes.ToTimestamp(val) }
