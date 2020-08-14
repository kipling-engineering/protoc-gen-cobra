package yaml

import (
	"io"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
)

func init() {
	client.RegisterInputDecoder("yaml", decoderMaker)
	client.RegisterOutputEncoder("yaml", encoderMaker)
}

func decoderMaker(r io.Reader) iocodec.Decoder {
	return yaml.NewDecoder(r).Decode
}

func encoderMaker(w io.Writer) iocodec.Encoder {
	return func(v interface{}) error {
		// workaround: yaml encoder doesn't honor json tags so pre-preprocess with mapstructure first
		m := make(map[string]interface{})
		cfg := &mapstructure.DecoderConfig{Result: &m, TagName: "json"}
		if dec, err := mapstructure.NewDecoder(cfg); err != nil {
			return err
		} else if err := dec.Decode(v); err != nil {
			return err
		}

		e := yaml.NewEncoder(w)
		defer e.Close()
		return e.Encode(m)
	}
}
