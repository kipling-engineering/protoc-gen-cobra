package yaml

import (
	"io"

	"gopkg.in/yaml.v2"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
)

func init() {
	client.DefaultConfig.RegisterDecoder("yaml", decoderMaker)
	client.DefaultConfig.RegisterEncoder("yaml", encoderMaker)
}

func decoderMaker(r io.Reader) iocodec.Decoder {
	return yaml.NewDecoder(r).Decode
}

func encoderMaker(w io.Writer) iocodec.Encoder {
	e := yaml.NewEncoder(w)
	defer e.Close()
	return e.Encode
}
