package yaml

import (
	"io"

	"gopkg.in/yaml.v2"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
)

func init() {
	client.DefaultConfig.RegisterDecoder("yaml", func(r io.Reader) iocodec.Decoder { return yaml.NewDecoder(r).Decode })
	client.DefaultConfig.RegisterEncoder("yaml", func(w io.Writer) iocodec.Encoder { return yaml.NewEncoder(w).Encode })
}
