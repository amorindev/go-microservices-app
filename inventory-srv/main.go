package main

import (
	"log"
	"net"

	"github.com/amorindev/go-microservices-app/inventory-srv/handler"
	"github.com/amorindev/go-microservices-app/inventory-srv/repository/memory"
	"github.com/amorindev/go-microservices-app/inventory-srv/service"
	"google.golang.org/grpc"
)

// GRPC_ADDR
var grpcAddr = "localhost:3100"

func main() {
	grpcServer := grpc.NewServer()

	repo := memory.NewRepo()
	srv := service.NewSrv(repo)

	hander := handler.NewHandler(srv)
	hander.RegisterService(grpcServer)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer l.Close()

	log.Println("INVENTORY-SRV Grpc Server Started at", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal()
	}

}
