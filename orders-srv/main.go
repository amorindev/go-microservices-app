package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPC_ADDR
var grpcAddr = "localhost:3000"

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer l.Close()

	store := NewStore()
	srv := NewSrv(store)

	hander := NewHandler(srv)
	hander.RegisterService(grpcServer)

	srv.Create(context.Background())

	log.Println("Grpc Server Started at", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal()
	}
}
