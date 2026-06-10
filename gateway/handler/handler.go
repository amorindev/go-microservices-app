package handler

import (
	"net/http"

	categoriesPB "github.com/amorindev/go-microservices-app/shared/categories/api/grpc/gen"
	ordersPB "github.com/amorindev/go-microservices-app/shared/orders/api/grpc/gen"
)

type Handler struct {
	orderC    ordersPB.OrderServiceClient
	categoryC categoriesPB.CategoryServiceClient
}

func NewHandler(orderClient ordersPB.OrderServiceClient, catalogC categoriesPB.CategoryServiceClient) *Handler {
	return &Handler{
		orderC:    orderClient,
		categoryC: catalogC,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// Orders
	mux.HandleFunc("GET /api/customers/{userID}/orders", h.HandleCreateOrder)

	// Categories
	mux.HandleFunc("POST /api/categories", h.HandleCreateCategory)
	mux.HandleFunc("GET /api/categories", h.HandleGetAllCategory)
}
