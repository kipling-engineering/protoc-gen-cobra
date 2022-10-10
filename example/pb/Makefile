PROTO_SOURCES = $(wildcard *.proto)
GO_SOURCES = $(patsubst %.proto,%.pb.go,$(PROTO_SOURCES))
GRPC_SOURCES = $(patsubst %.proto,%_grpc.pb.go,$(PROTO_SOURCES))
COBRA_SOURCES = $(patsubst %.proto,%.cobra.pb.go,$(PROTO_SOURCES))

all: $(GO_SOURCES) $(GRPC_SOURCES) $(COBRA_SOURCES)

$(GO_SOURCES) $(GRPC_SOURCES) $(COBRA_SOURCES):
	protoc \
		-I. \
		--go_out=. \
		--go-grpc_out=. \
		--cobra_out=. \
		$(PROTO_SOURCES)

clean:
	rm -f $(GO_SOURCES) $(GRPC_SOURCES) $(COBRA_SOURCES)

deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/NathanBaulch/protoc-gen-cobra@latest
