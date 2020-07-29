package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/pflag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/NathanBaulch/protoc-gen-cobra/iocodec"
)

type Config struct {
	ServerAddr     string
	RequestFile    string
	RequestFormat  string
	ResponseFormat string
	Timeout        time.Duration

	TLS                bool
	ServerName         string
	InsecureSkipVerify bool
	CACertFile         string
	CertFile           string
	KeyFile            string

	flags       []func(*pflag.FlagSet)
	dialOptions []func(context.Context, *[]grpc.DialOption) error
	decoders    map[string]iocodec.DecoderMaker
	encoders    map[string]iocodec.EncoderMaker
}

var DefaultConfig = &Config{
	ServerAddr:     "localhost:8080",
	RequestFormat:  "json",
	ResponseFormat: "json",
	Timeout:        10 * time.Second,

	decoders: map[string]iocodec.DecoderMaker{
		"json": func(r io.Reader) iocodec.Decoder { return json.NewDecoder(r).Decode },
		"xml":  func(r io.Reader) iocodec.Decoder { return xml.NewDecoder(r).Decode },
	},
	encoders: map[string]iocodec.EncoderMaker{
		"json":       func(w io.Writer) iocodec.Encoder { return json.NewEncoder(w).Encode },
		"prettyjson": func(w io.Writer) iocodec.Encoder { e := json.NewEncoder(w); e.SetIndent("", "  "); return e.Encode },
		"xml":        func(w io.Writer) iocodec.Encoder { return xml.NewEncoder(w).Encode },
		"prettyxml":  func(w io.Writer) iocodec.Encoder { e := xml.NewEncoder(w); e.Indent("", "  "); return e.Encode },
	},
}

func (c Config) Clone() *Config {
	d := make(map[string]iocodec.DecoderMaker, len(c.decoders))
	for k, v := range c.decoders {
		d[k] = v
	}
	c.decoders = d
	e := make(map[string]iocodec.EncoderMaker, len(c.encoders))
	for k, v := range c.encoders {
		e[k] = v
	}
	c.encoders = e
	return &c
}

func (c *Config) RegisterFlags(f func(fs *pflag.FlagSet)) {
	c.flags = append(c.flags, f)
}

func (c *Config) RegisterDialOptions(f func(context.Context, *[]grpc.DialOption) error) {
	c.dialOptions = append(c.dialOptions, f)
}

func (c *Config) RegisterDecoder(format string, maker iocodec.DecoderMaker) {
	c.decoders[format] = maker
}

func (c *Config) RegisterEncoder(format string, maker iocodec.EncoderMaker) {
	c.encoders[format] = maker
}

func (c *Config) BindFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&c.ServerAddr, "server-addr", "s", c.ServerAddr, "server address in the form host:port")
	fs.StringVarP(&c.RequestFile, "request-file", "f", c.RequestFile, "client request file; use \"-\" for stdin")
	fs.StringVarP(&c.RequestFormat, "request-format", "i", c.RequestFormat, "request format ("+strings.Join(c.decoderFormats(), ", ")+")")
	fs.StringVarP(&c.ResponseFormat, "response-format", "o", c.ResponseFormat, "response format ("+strings.Join(c.encoderFormats(), ", ")+")")
	fs.DurationVar(&c.Timeout, "timeout", c.Timeout, "client connection timeout")
	fs.BoolVar(&c.TLS, "tls", c.TLS, "enable TLS")
	fs.StringVar(&c.ServerName, "tls-server-name", c.ServerName, "TLS server name override")
	fs.BoolVar(&c.InsecureSkipVerify, "tls-insecure-skip-verify", c.InsecureSkipVerify, "INSECURE: skip TLS checks")
	fs.StringVar(&c.CACertFile, "tls-ca-cert-file", c.CACertFile, "CA certificate file")
	fs.StringVar(&c.CertFile, "tls-cert-file", c.CertFile, "client certificate file")
	fs.StringVar(&c.KeyFile, "tls-key-file", c.KeyFile, "client key file")

	for _, h := range c.flags {
		h(fs)
	}
}

func (c *Config) decoderFormats() []string {
	f := make([]string, len(c.decoders))
	i := 0
	for k := range c.decoders {
		f[i] = k
		i++
	}
	return f
}

func (c *Config) encoderFormats() []string {
	f := make([]string, len(c.encoders))
	i := 0
	for k := range c.encoders {
		f[i] = k
		i++
	}
	return f
}

type Dialer struct{ *Config }

func (d *Dialer) RoundTrip(ctx context.Context, fn func(grpc.ClientConnInterface, iocodec.Decoder, iocodec.Encoder) error) error {
	var err error
	var in iocodec.Decoder
	if in, err = d.makeDecoder(); err != nil {
		return err
	}
	var out iocodec.Encoder
	if out, err = d.makeEncoder(); err != nil {
		return err
	}

	opts := []grpc.DialOption{grpc.WithBlock()}
	if err := d.dialOpts(ctx, &opts); err != nil {
		return err
	}

	if d.Timeout > 0 {
		var done context.CancelFunc
		ctx, done = context.WithTimeout(ctx, d.Timeout)
		defer done()
	}

	cc, err := grpc.DialContext(ctx, d.ServerAddr, opts...)
	if err != nil {
		if err == context.DeadlineExceeded {
			return fmt.Errorf("timeout dialing server: %s", d.ServerAddr)
		}
		return err
	}
	defer cc.Close()

	return fn(cc, in, out)
}

func (d *Dialer) makeDecoder() (iocodec.Decoder, error) {
	var r io.Reader
	if stat, _ := os.Stdin.Stat(); (stat.Mode()&os.ModeCharDevice) == 0 || d.RequestFile == "-" {
		r = os.Stdin
	} else if d.RequestFile != "" {
		f, err := os.Open(d.RequestFile)
		if err != nil {
			return nil, fmt.Errorf("request file: %v", err)
		}
		defer f.Close()
		if ext := strings.TrimLeft(filepath.Ext(d.RequestFile), "."); ext != "" {
			if m, ok := d.decoders[ext]; ok {
				return m(f), nil
			}
		}
		r = f
	} else {
		r = nil
	}

	if r == nil || d.RequestFormat == "" {
		return iocodec.NoOp, nil
	}
	if m, ok := d.decoders[d.RequestFormat]; !ok {
		return nil, fmt.Errorf("unknown request format: %s", d.RequestFormat)
	} else {
		return m(r), nil
	}
}

func (d *Dialer) makeEncoder() (iocodec.Encoder, error) {
	if d.ResponseFormat == "" {
		return iocodec.NoOp, nil
	}
	if m, ok := d.encoders[d.ResponseFormat]; !ok {
		return nil, fmt.Errorf("unknown response format: %s", d.ResponseFormat)
	} else {
		return m(os.Stdout), nil
	}
}

func (d *Dialer) dialOpts(ctx context.Context, opts *[]grpc.DialOption) error {
	if d.TLS {
		tlsConfig := &tls.Config{InsecureSkipVerify: d.InsecureSkipVerify}
		if d.CACertFile != "" {
			caCert, err := ioutil.ReadFile(d.CACertFile)
			if err != nil {
				return fmt.Errorf("ca cert: %v", err)
			}
			certPool := x509.NewCertPool()
			certPool.AppendCertsFromPEM(caCert)
			tlsConfig.RootCAs = certPool
		}
		if d.CertFile != "" {
			if d.KeyFile == "" {
				return fmt.Errorf("key file not specified")
			}
			pair, err := tls.LoadX509KeyPair(d.CertFile, d.KeyFile)
			if err != nil {
				return fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}
		if d.ServerName != "" {
			tlsConfig.ServerName = d.ServerName
		} else {
			addr, _, _ := net.SplitHostPort(d.ServerAddr)
			tlsConfig.ServerName = addr
		}
		cred := credentials.NewTLS(tlsConfig)
		*opts = append(*opts, grpc.WithTransportCredentials(cred))
	} else {
		*opts = append(*opts, grpc.WithInsecure())
	}

	for _, h := range d.dialOptions {
		if err := h(ctx, opts); err != nil {
			return err
		}
	}

	return nil
}
