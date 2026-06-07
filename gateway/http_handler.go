package main

import (
	"errors"
	"net/http"

	"github.com/amorindev/go-microservices-app/shared"
	ordersPB "github.com/amorindev/go-microservices-app/shared/api/proto/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	orderC ordersPB.OrderServiceClient
}

func NewHandler(client ordersPB.OrderServiceClient) *handler {
	return &handler{
		orderC: client,
	}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*ordersPB.ItemsWithQuantity
	if err := shared.ReadJSON(r, &items); err != nil {
		shared.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		shared.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.orderC.Create(r.Context(), &ordersPB.CreateOrderReq{
		CustomerID: customerID,
		Items:      items,
	})
	rStatus := status.Convert(err)
	if err != nil {
		if rStatus.Code() != codes.InvalidArgument {
			shared.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}
		shared.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*ordersPB.ItemsWithQuantity) error {
	if len(items) == 0 {
		return shared.ErrNoItems
	}

	for _, i := range items {
		if i.ID == "" {
			return errors.New("item ID is required")
		}
		if i.Quantity <= 0 {
			return errors.New("items must be have a valid quantity")
		}
	}
	return nil
}
