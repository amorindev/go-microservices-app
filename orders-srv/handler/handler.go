package handler

import (
	"github.com/amorindev/go-microservices-app/orders-srv/port"
	ordersPB "github.com/amorindev/go-microservices-app/shared/orders/api/grpc/gen"
	"google.golang.org/grpc"
)

type Handler struct {
	ordersPB.UnimplementedOrderServiceServer
	service port.OrdersSrv
}

func NewHandler(ordersSrv port.OrdersSrv) *Handler {
	return &Handler{
		service: ordersSrv,
	}
}

func (h *Handler) RegisterService(grpcServer *grpc.Server) {
	ordersPB.RegisterOrderServiceServer(grpcServer, h)
}