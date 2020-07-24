package flag

import (
	"encoding/base64"
	"strings"

	"github.com/spf13/pflag"
)

type bytesBase64SliceValue struct {
	value   *[][]byte
	changed bool
}

func BytesBase64SliceVar(fs *pflag.FlagSet, p *[][]byte, name, usage string) {
	fs.Var(&bytesBase64SliceValue{value: p}, name, usage)
}

func (s *bytesBase64SliceValue) Set(val string) error {
	ss := strings.Split(val, ",")
	out := make([][]byte, len(ss))
	for i, v := range ss {
		var err error
		if out[i], err = s.fromString(v); err != nil {
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

func (s *bytesBase64SliceValue) Type() string { return "bytesBase64Slice" }

func (s *bytesBase64SliceValue) String() string { return "[]" }

func (s *bytesBase64SliceValue) Append(val string) error {
	d, err := s.fromString(val)
	if err != nil {
		return err
	}
	*s.value = append(*s.value, d)
	return nil
}

func (s *bytesBase64SliceValue) Replace(val []string) error {
	out := make([][]byte, len(val))
	for i, v := range val {
		var err error
		out[i], err = s.fromString(v)
		if err != nil {
			return err
		}
	}
	*s.value = out
	return nil
}

func (s *bytesBase64SliceValue) GetSlice() []string {
	out := make([]string, len(*s.value))
	for i, d := range *s.value {
		out[i] = s.toString(d)
	}
	return out
}

func (s *bytesBase64SliceValue) fromString(val string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(val)
}

func (s *bytesBase64SliceValue) toString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
