package handler

import (
	"context"

	pb "hello_rpc/api"
)

type GreeterHandler struct {
	pb.UnimplementedHelloServer
}

func (GreeterHandler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + req.GetName(),
	}, nil
}
