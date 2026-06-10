package handler

import (
	"github.com/amorindev/go-microservices-app/inventory-srv/port"
	inventoryPB "github.com/amorindev/go-microservices-app/shared/inventory/api/grpc/gen"
	"google.golang.org/grpc"
)

type Handler struct {
	inventoryPB.UnimplementedInventoryServiceServer
	service port.InventorySrv
}

func NewHandler(inventorySrv port.InventorySrv) *Handler {
	return &Handler{
		service: inventorySrv,
	}
}

func (h *Handler) RegisterService(grpcServer *grpc.Server) {
	inventoryPB.RegisterInventoryServiceServer(grpcServer, h)
}