package main

import (
	"cmp"
	"log"
	"net/http"
	"os"

	"github.com/amorindev/go-microservices-app/gateway/handler"
	categoriesPB "github.com/amorindev/go-microservices-app/shared/categories/api/grpc/gen"
	ordersPB "github.com/amorindev/go-microservices-app/shared/orders/api/grpc/gen"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	orderServiceAddr = "localhost:3000"
	catalogServiceAddr = "localhost:3200"
)

func main() {

	// Order service
	orderConn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer orderConn.Close()

	log.Println("GATEWAY - Dialing orders service at", orderServiceAddr)

	orderC := ordersPB.NewOrderServiceClient(orderConn)


	// Category service
	categoryConn, err := grpc.NewClient(catalogServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer categoryConn.Close()

	log.Println("GATEWAY - Dialing catalog service at", catalogServiceAddr)

	catalogC := categoriesPB.NewCategoryServiceClient(categoryConn)

	err = godotenv.Load("./gateway/.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	hsp := cmp.Or(os.Getenv("HTTP_SERVER_PORT"), "8000")
	hsp = ":" + hsp

	mux := http.NewServeMux()

	handler := handler.NewHandler(orderC,catalogC)
	handler.RegisterRoutes(mux)

	log.Printf("GATEWAY - Starting HTTP server at %s", hsp)

	if err := http.ListenAndServe(hsp, mux); err != nil {
		log.Fatal("Failed to start http server")
	}

}
