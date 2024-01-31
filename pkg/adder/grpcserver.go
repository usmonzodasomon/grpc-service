package adder

import (
	"context"

	adder "github.com/usmonzodasomon/grpc-service/pkg/proto"
)

type GRPCServer struct {
	adder.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *adder.AddRequest) (*adder.AddResponse, error) {
	return &adder.AddResponse{
		Result: req.GetX() + req.GetY(),
	}, nil
}
