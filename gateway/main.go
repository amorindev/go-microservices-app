package main

import (
	"cmp"
	"log"
	"net/http"
	"os"

	ordersPB "github.com/amorindev/go-microservices-app/shared/api/proto/gen"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	orderServiceAddr = "localhost:3000"
)

func main() {

	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at", orderServiceAddr)

	c := ordersPB.NewOrderServiceClient(conn)

	err = godotenv.Load("./gateway/.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	hsp := cmp.Or(os.Getenv("HTTP_SERVER_PORT"), "8000")
	hsp = ":" + hsp

	mux := http.NewServeMux()

	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", hsp)

	if err := http.ListenAndServe(hsp, mux); err != nil {
		log.Fatal("Failed to start http server")
	}

}
