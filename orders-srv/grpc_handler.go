package main

import (
	"context"
	"log"

	ordersPB "github.com/amorindev/go-microservices-app/shared/api/proto/gen"
	"google.golang.org/grpc"
)

type Handler struct {
	ordersPB.UnimplementedOrderServiceServer
	service OrdersSrv
}

func NewHandler(ordersSrv OrdersSrv) *Handler {
	return &Handler{
		service: ordersSrv,
	}
}

func (h *Handler) RegisterService(grpcServer *grpc.Server) {
	ordersPB.RegisterOrderServiceServer(grpcServer, h)
}

func (h *Handler) Create(ctx context.Context, req *ordersPB.CreateOrderReq) (*ordersPB.Order, error) {
	//return nil, fmt.Errorf("some errors!!")
	log.Printf("New order received!!: %v", req)
	o := &ordersPB.Order{
		Id: "42",
	}

	return o, nil
}
