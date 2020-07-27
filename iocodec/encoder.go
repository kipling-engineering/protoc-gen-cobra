package iocodec

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

var encoders = map[string]EncoderMaker{
	"noop":       func(w io.Writer) Encoder { return func(interface{}) error { return nil } },
	"xml":        func(w io.Writer) Encoder { return xml.NewEncoder(w).Encode },
	"prettyxml":  func(w io.Writer) Encoder { e := xml.NewEncoder(w); e.Indent("", "  "); return e.Encode },
	"json":       func(w io.Writer) Encoder { return json.NewEncoder(w).Encode },
	"prettyjson": func(w io.Writer) Encoder { e := json.NewEncoder(w); e.SetIndent("", "  "); return e.Encode },
}

type (
	EncoderMaker func(io.Writer) Encoder
	Encoder      func(interface{}) error
)

func RegisterEncoder(format string, maker EncoderMaker) {
	encoders[format] = maker
}

func EncoderFormats() []string {
	f := make([]string, len(encoders))
	i := 0
	for k := range encoders {
		f[i] = k
		i++
	}
	return f
}

func MakeEncoder(format string, r io.Writer) Encoder {
	if format == "" {
		format = "noop"
	}
	if m, ok := encoders[format]; !ok {
		return nil
	} else {
		return m(r)
	}
}
