package iocodec

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

var NoOp = func(interface{}) error { return nil }

type (
	DecoderMaker func(io.Reader) Decoder
	Decoder      func(interface{}) error
	EncoderMaker func(io.Writer) Encoder
	Encoder      func(interface{}) error
)

func JSONDecoderMaker(r io.Reader) Decoder {
	return json.NewDecoder(r).Decode
}

func JSONEncoderMaker(pretty bool) EncoderMaker {
	return func(w io.Writer) Encoder {
		e := json.NewEncoder(w)
		if pretty {
			e.SetIndent("", "  ")
		}
		return e.Encode
	}
}

func XMLDecoderMaker(r io.Reader) Decoder {
	return xml.NewDecoder(r).Decode
}

func XMLEncoderMaker(pretty bool) EncoderMaker {
	return func(w io.Writer) Encoder {
		return func(v interface{}) error {
			e := xml.NewEncoder(w)
			defer e.Flush()
			if pretty {
				e.Indent("", "  ")
			}
			if err := e.Encode(v); err != nil {
				return err
			}
			_, err := w.Write([]byte("\n"))
			return err
		}
	}
}
