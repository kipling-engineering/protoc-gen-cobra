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
		if out[i], err = base64.RawStdEncoding.DecodeString(v); err != nil {
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

func (*bytesBase64SliceValue) Type() string { return "bytesBase64Slice" }

func (*bytesBase64SliceValue) String() string { return "[]" }
