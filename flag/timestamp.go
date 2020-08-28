package flag

import (
	"strings"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/pflag"

	"github.com/NathanBaulch/protoc-gen-cobra/ptypes"
)

func TimestampVar(fs *pflag.FlagSet, p **timestamp.Timestamp, name, usage string) {
	v := fs.String(name, "", usage)
	hook := func() error {
		if d, err := ptypes.ToTimestamp(v); err != nil {
			return err
		} else {
			*p = d
			return nil
		}
	}
	WithPostSetHookE(fs, name, hook)
}

type timestampSliceValue struct {
	value   *[]*timestamp.Timestamp
	changed bool
}

func TimestampSliceVar(fs *pflag.FlagSet, p *[]*timestamp.Timestamp, name, usage string) {
	fs.Var(&timestampSliceValue{value: p}, name, usage)
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
	if !s.changed {
		*s.value = out
		s.changed = true
	} else {
		*s.value = append(*s.value, out...)
	}
	return nil
}

func (*timestampSliceValue) Type() string { return "timestampSlice" }

func (*timestampSliceValue) String() string { return "[]" }

func ParseTimestamp(val string) (interface{}, error) { return ptypes.ToTimestamp(val) }
