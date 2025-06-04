PROTO_SOURCES = $(wildcard *.proto)
GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
GRPC_SOURCES = $(patsubst %.proto,%_grpc.pb.go,$(PROTO_SOURCES))
COBRA_SOURCES = $(patsubst %.proto,%.cobra.pb.go,$(PROTO_SOURCES))

# Determine GOBIN, assuming GOPATH/bin or HOME/go/bin
GOBIN := $(shell go env GOPATH)/bin
ifeq ($(wildcard $(GOBIN)/protoc-gen-go),)
    GOBIN := $(HOME)/go/bin
endif

all:
	@echo "Using GOBIN: $(GOBIN)"
	protoc \
		-I. \
		-I/tmp/protoc/include \
		--plugin=protoc-gen-go=$(GOBIN)/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(GOBIN)/protoc-gen-go-grpc \
		--plugin=protoc-gen-cobra=$(GOBIN)/protoc-gen-cobra \
		--go_out=. \
		--go-grpc_out=. \
		--cobra_out=. \
		$(PROTO_SOURCES)

clean:
	rm -f $(GO_SOURCES) $(GRPC_SOURCES) $(COBRA_SOURCES)

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/NathanBaulch/protoc-gen-cobra@v1.2.1
