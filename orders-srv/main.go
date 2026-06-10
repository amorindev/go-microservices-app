package main

import (
	"log"
	"net"

	"github.com/amorindev/go-microservices-app/orders-srv/handler"
	"github.com/amorindev/go-microservices-app/orders-srv/repository/memory"
	"github.com/amorindev/go-microservices-app/orders-srv/service"
	"google.golang.org/grpc"
)

// GRPC_ADDR
var grpcAddr = "localhost:3000"

func main() {

	grpcServer := grpc.NewServer()

	store := memory.NewStore()
	srv := service.NewSrv(store)

	handler := handler.NewHandler(srv)
	handler.RegisterService(grpcServer)

	
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	defer l.Close()

	log.Println("ORDERS-SRV Grpc Server Started at", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal()
	}
}
