// Copyright 2016 The protoc-gen-cobra authors. All rights reserved.

// protoc-gen-cobra is a plugin for the Google protocol buffer compiler to
// generate Go code to be used for building command line clients using cobra.
// Run it by building this program and putting it in your path with the name
// 	protoc-gen-cobra
// That word 'cobra' at the end becomes part of the option string set for the
// protocol compiler, so once the protocol compiler (protoc) is installed
// you can run
// 	protoc --cobra_out=output_directory input_directory/file.proto
// to generate Go bindings for the protocol defined by file.proto.
// With that input, the output will be written to
// 	output_directory/file.pb.go
// Use it combined with the go and go-grpc outputs
//	protoc --go_out=. --go-grpc_out=. --cobra_out=. *.proto
//
// The generated code is documented in the package comment for
// the library.
//
// See the README and documentation for protocol buffers to learn more:
// 	https://developers.google.com/protocol-buffers/

package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				if err := genFile(gen, f); err != nil {
					return err
				}
			}
		}
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return nil
	})
}
