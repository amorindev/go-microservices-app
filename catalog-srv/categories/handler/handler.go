package handler

import (
	"github.com/amorindev/go-microservices-app/catalog-srv/categories/port"
	categoriesPB "github.com/amorindev/go-microservices-app/shared/categories/api/grpc/gen"
	"google.golang.org/grpc"
)

type Handler struct {
	categoriesPB.UnimplementedCategoryServiceServer
	service port.CategoryRepo
}

func NewHandler(categorySrv port.CategorySrv) *Handler {
	return &Handler{
		service: categorySrv,
	}
}

func (h *Handler) RegisterService(grpcServer *grpc.Server) {
	categoriesPB.RegisterCategoryServiceServer(grpcServer, h)
}