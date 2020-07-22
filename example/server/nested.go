package main

import (
	"golang.org/x/net/context"

	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"
)

type Nested struct {
	pb.UnimplementedNestedServer
}

func NewNested() *Nested {
	return &Nested{}
}

func (*Nested) Get(_ context.Context, req *pb.NestedRequest) (*pb.NestedResponse, error) {
	return &pb.NestedResponse{
		Return: req.Inner.Value + req.TopLevel.Value,
	}, nil
}

func (*Nested) GetDeeplyNested(_ context.Context, req *pb.DeeplyNested) (*pb.NestedResponse, error) {
	return &pb.NestedResponse{
		Return: req.L0.L1.L2.L3,
	}, nil
}
