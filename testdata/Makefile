# Determine GOBIN, assuming GOPATH/bin or HOME/go/bin
GOBIN := $(shell go env GOPATH)/bin
ifeq ($(wildcard $(GOBIN)/protoc-gen-go),)
    GOBIN := $(HOME)/go/bin
endif

# All .proto files in pb/
PB_PROTO_SOURCES := $(wildcard pb/*.proto)

# Define generated file names based on sources
PB_GO_GENERATED := $(patsubst pb/%.proto,pb/%.pb.go,$(PB_PROTO_SOURCES))
PB_GRPC_GENERATED := $(patsubst pb/%.proto,pb/%_grpc.pb.go,$(PB_PROTO_SOURCES))
PB_COBRA_GENERATED := $(patsubst pb/%.proto,pb/%.cobra.pb.go,$(PB_PROTO_SOURCES))

# Minimal all target for now, focusing on pb files
.PHONY: all clean deps

all: $(PB_GO_GENERATED) $(PB_GRPC_GENERATED) $(PB_COBRA_GENERATED)
	@echo "Selected files generated."

# Pattern rule for pb/*.proto files
# This rule is run from the 'testdata' directory (due to make -C testdata)
# Input file ($<) is relative to 'testdata', e.g., pb/fieldmask_test.proto
# Output files (targets) are relative to 'testdata', e.g., pb/fieldmask_test.pb.go
# --*_out=. means output root is 'testdata'
# go_package="./pb" in proto ensures files are placed in 'testdata/pb/'
pb/%.pb.go pb/%_grpc.pb.go pb/%.cobra.pb.go: pb/%.proto
	@echo "Generating files for $< (go, grpc, cobra)..."
	@echo "Using GOBIN: $(GOBIN)"
	protoc \
		-I . \
		--plugin=protoc-gen-go=$(GOBIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(GOBIN)/protoc-gen-go-grpc \
		--plugin=protoc-gen-cobra=$(GOBIN)/protoc-gen-cobra \
		--go_out=. \
		--go-grpc_out=. \
		--cobra_out=. \
		$<

clean:
	rm -f $(PB_GO_GENERATED) $(PB_GRPC_GENERATED) $(PB_COBRA_GENERATED)
	# Add other clean rules if necessary for other proto directories from original Makefile eventually
	@echo "Cleaned generated pb/ files."

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/NathanBaulch/protoc-gen-cobra@latest
