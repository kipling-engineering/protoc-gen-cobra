package flag

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

var nonWordChars = regexp.MustCompile(`[^\w]`)

func SetFlagsFromEnv(fs *pflag.FlagSet, prefixes ...string) (err error) {
	parts := make([]string, 0, len(prefixes)+1)
	for _, prefix := range prefixes {
		if prefix != "" {
			parts = append(parts, format(prefix))
		}
	}

	fs.VisitAll(func(f *pflag.Flag) {
		if err != nil || f.Changed {
			return
		}
		name := strings.Join(append(parts, format(f.Name)), "_")
		if val := os.Getenv(name); val != "" {
			if err = f.Value.Set(val); err != nil {
				err = fmt.Errorf("environment variable %s: %v", name, err)
			}
		}
	})

	return
}

func format(name string) string {
	return strings.ToUpper(nonWordChars.ReplaceAllString(name, "_"))
}
