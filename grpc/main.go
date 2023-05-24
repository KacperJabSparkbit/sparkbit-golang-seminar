package main

import (
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	protos "webservices/grpc/protos/proto/product"
	"webservices/grpc/server"
)

func main() {
	log := hclog.Default()
	log.Info("Starting the server...")

	gs := grpc.NewServer()

	ps := server.NewProductServer(log)

	protos.RegisterProductServer(gs, ps)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		panic(err)
	}

	err = gs.Serve(l)
	if err != nil {
		return
	}
}
