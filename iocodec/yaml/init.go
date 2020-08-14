package yaml

import (
	"io"

	"github.com/fatih/structs"
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
	return func(v interface{}) error {
		// workaround: yaml decoder doesn't honor json tags so decode to map and populate the target struct with mapstructure
		m := make(map[string]interface{})
		if err := yaml.NewDecoder(r).Decode(&m); err != nil {
			return err
		}
		config := &mapstructure.DecoderConfig{Result: v, TagName: "json"}
		if decoder, err := mapstructure.NewDecoder(config); err != nil {
			return err
		} else {
			return decoder.Decode(m)
		}
	}
}

func encoderMaker(w io.Writer) iocodec.Encoder {
	return func(v interface{}) error {
		// workaround: yaml encoder doesn't honor json tags so encode to map first with structs 
		s := structs.New(v)
		s.TagName = "json"
		v = s.Map()

		e := yaml.NewEncoder(w)
		defer e.Close()
		return e.Encode(v)
	}
}
