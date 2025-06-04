module github.com/kipling-engineering/protoc-gen-cobra

go 1.21 // Reverted to 1.21

require (
	github.com/google/go-cmp v0.6.0 // Reverted or kept older
	github.com/iancoleman/strcase v0.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/spf13/cast v1.5.0
	github.com/spf13/cobra v1.6.1 // Reverted or kept older
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	github.com/stretchr/testify v1.8.1
	golang.org/x/net v0.22.0 // Reverted or kept older
	golang.org/x/oauth2 v0.18.0 // Reverted or kept older
	google.golang.org/grpc v1.64.0 // Trying a newer version known to be stable with Go 1.21
	google.golang.org/protobuf v1.33.0 // Version compatible with grpc v1.60.x
	gopkg.in/yaml.v3 v3.0.1
)

// Indirect dependencies will be re-evaluated by 'go mod tidy'
// Keeping the old list for now, tidy will fix it.
require (
	cloud.google.com/go/compute v1.25.1 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/afero v1.9.4 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)

require github.com/NathanBaulch/protoc-gen-cobra v1.2.1

replace github.com/NathanBaulch/protoc-gen-cobra => ./

require google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
