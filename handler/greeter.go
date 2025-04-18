package handler

import (
	"context"

	pb "hello_rpc/api"
)

type GreeterHandler struct {
	pb.UnimplementedHelloServer
	pb.UnimplementedTestServer
}

func (GreeterHandler) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello rpc abc " + req.GetName(),
	}, nil
}

func (GreeterHandler) Ping(ctx context.Context, req *pb.PingReq) (*pb.PingResp, error) {
	return &pb.PingResp{
		Resp: "pong",
	}, nil
}
