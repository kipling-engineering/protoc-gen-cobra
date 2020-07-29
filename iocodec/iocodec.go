package iocodec

import "io"

var NoOp = func(interface{}) error { return nil }

type (
	DecoderMaker func(io.Reader) Decoder
	Decoder      func(interface{}) error
	EncoderMaker func(io.Writer) Encoder
	Encoder      func(interface{}) error
)
