package oauth

import (
	"context"

	"github.com/spf13/pflag"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/NathanBaulch/protoc-gen-cobra/client"
)

var Config = &config{
	TokenType: "Bearer",
}

type config struct {
	AccessToken string
	TokenType   string
}

func init() {
	client.RegisterFlagBinder(func(fs *pflag.FlagSet) {
		fs.StringVar(&Config.AccessToken, "auth-access-token", Config.AccessToken, "authorization access token")
		fs.StringVar(&Config.TokenType, "auth-token-type", Config.TokenType, "authorization token type")
	})

	client.RegisterPreDialer(func(_ context.Context, opts *[]grpc.DialOption) error {
		cfg := Config

		if cfg.AccessToken != "" {
			cred := oauth.NewOauthAccess(&oauth2.Token{
				AccessToken: cfg.AccessToken,
				TokenType:   cfg.TokenType,
			})
			*opts = append(*opts, grpc.WithPerRPCCredentials(cred))
		}

		return nil
	})
}
