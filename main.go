package main

import (
	"github.com/dreamilk/rpc_server/server"

	pb "hello_rpc/api"
	"hello_rpc/handler"
)

func main() {
	server.RegisterService(&pb.Hello_ServiceDesc, &handler.GreeterHandler{})

	if err := server.Serve(); err != nil {
		return
	}
}
