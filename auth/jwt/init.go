package jwt

import (
	"context"
	"fmt"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
)

var Config = &config{}

type config struct {
	JWTKey     string
	JWTKeyFile string
}

func init() {
	client.RegisterFlagBinder(func(fs *pflag.FlagSet) {
		fs.StringVar(&Config.JWTKey, "jwt-key", Config.JWTKey, "JWT key")
		fs.StringVar(&Config.JWTKeyFile, "jwt-key-file", Config.JWTKeyFile, "JWT key file")
	})

	client.RegisterPreDialer(func(_ context.Context, opts *[]grpc.DialOption) error {
		cfg := Config

		if cfg.JWTKey != "" {
			cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.JWTKey))
			if err != nil {
				return fmt.Errorf("jwt key: %v", err)
			}
			*opts = append(*opts, grpc.WithPerRPCCredentials(cred))
		}
		if cfg.JWTKeyFile != "" {
			cred, err := oauth.NewJWTAccessFromFile(cfg.JWTKeyFile)
			if err != nil {
				return fmt.Errorf("jwt key file: %v", err)
			}
			*opts = append(*opts, grpc.WithPerRPCCredentials(cred))
		}

		return nil
	})
}
