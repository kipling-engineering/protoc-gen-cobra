package flag

import (
	"fmt"
	"os"
	"strings"

	"github.com/NathanBaulch/protoc-gen-cobra/naming"
	"github.com/spf13/pflag"
)

func SetFlagsFromEnv(fs *pflag.FlagSet, namer naming.Namer, prefixes ...string) (err error) {
	parts := make([]string, 0, len(prefixes)+1)
	for _, prefix := range prefixes {
		if prefix != "" {
			parts = append(parts, namer(prefix))
		}
	}

	fs.VisitAll(func(f *pflag.Flag) {
		if err != nil || f.Changed {
			return
		}
		name := strings.Join(append(parts, namer(f.Name)), "_")
		if val := os.Getenv(name); val != "" {
			if err = f.Value.Set(val); err != nil {
				err = fmt.Errorf("environment variable %s: %v", name, err)
			}
		}
	})

	return
}
