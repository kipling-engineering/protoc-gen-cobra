package iocodec

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

var decoders = map[string]DecoderMaker{
	"noop": func(r io.Reader) Decoder { return func(interface{}) error { return nil } },
	"xml":  func(r io.Reader) Decoder { return xml.NewDecoder(r).Decode },
	"json": func(r io.Reader) Decoder { return json.NewDecoder(r).Decode },
}

type (
	DecoderMaker func(io.Reader) Decoder
	Decoder      func(interface{}) error
)

func RegisterDecoder(format string, maker DecoderMaker) {
	decoders[format] = maker
}

func DecoderFormats() []string {
	f := make([]string, len(decoders))
	i := 0
	for k := range decoders {
		f[i] = k
		i++
	}
	return f
}

func MakeDecoder(format string, r io.Reader) Decoder {
	if format == "" {
		format = "noop"
	}
	if m, ok := decoders[format]; !ok {
		return nil
	} else {
		return m(r)
	}
}
