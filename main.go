package main

import (
	"github.com/dreamilk/rpc_server/server"

	pb "login/api"
	"login/handler"
)

func main() {
	server.RegisterService(&pb.Hello_ServiceDesc, &handler.GreeterHandler{})

	if err := server.Serve(); err != nil {
		return
	}
}
