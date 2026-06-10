package main

import (
	"log"
	"net"

	categoryHandler "github.com/amorindev/go-microservices-app/catalog-srv/categories/handler"
	categoryMemory "github.com/amorindev/go-microservices-app/catalog-srv/categories/repository/memory"
	categoryService "github.com/amorindev/go-microservices-app/catalog-srv/categories/service"
	"google.golang.org/grpc"
)

// GRPC_ADDR
var grpcAddr = "localhost:3200"

func main() {

	grpcServer := grpc.NewServer()

	categoryRepo := categoryMemory.NewStore()
	categorySrv := categoryService.NewSrv(categoryRepo)

	handler := categoryHandler.NewHandler(categorySrv)
	handler.RegisterService(grpcServer)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer l.Close()

	log.Println("CATALOG-SRV Grpc Server Started at", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal()
	}
}
