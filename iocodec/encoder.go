package iocodec

import (
	"encoding/json"
	"encoding/xml"
	"io"

	"gopkg.in/yaml.v2"
)

// DefaultEncoders contains the default list of encoders per MIME type.
var DefaultEncoders = EncoderGroup{
	"xml":        EncoderMakerFunc(func(w io.Writer) Encoder { return xml.NewEncoder(w) }),
	"prettyxml":  EncoderMakerFunc(func(w io.Writer) Encoder { e := xml.NewEncoder(w); e.Indent("", "\t"); return e }),
	"json":       EncoderMakerFunc(func(w io.Writer) Encoder { return json.NewEncoder(w) }),
	"prettyjson": EncoderMakerFunc(func(w io.Writer) Encoder { e := json.NewEncoder(w); e.SetIndent("", "\t"); return e }),
	"yaml":       EncoderMakerFunc(func(w io.Writer) Encoder { return yaml.NewEncoder(w) }),
}

type (
	// An Encoder encodes data from v.
	Encoder interface {
		Encode(v interface{}) error
	}

	// An EncoderGroup maps MIME types to EncoderMakers.
	EncoderGroup map[string]EncoderMaker

	// An EncoderMaker creates and returns a new Encoder.
	EncoderMaker interface {
		NewEncoder(w io.Writer) Encoder
	}

	// EncoderMakerFunc is an adapter for creating EncoderMakers from functions.
	EncoderMakerFunc func(w io.Writer) Encoder
)

// NewEncoder implements the EncoderMaker interface.
func (f EncoderMakerFunc) NewEncoder(w io.Writer) Encoder {
	return f(w)
}
