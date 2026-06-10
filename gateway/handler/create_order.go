package handler

import (
	"net/http"

	"github.com/amorindev/go-microservices-app/shared"
	ordersPB "github.com/amorindev/go-microservices-app/shared/orders/api/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (h *Handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	
	var req *ordersPB.CreateOrderReq
	if err := shared.ReadJSON(r, &req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	req.UserId = userID

	o, err := h.orderC.Create(r.Context(), req)
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


