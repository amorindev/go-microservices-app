package handler

import (
	"net/http"

	"github.com/amorindev/go-microservices-app/shared"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) HandleGetAllCategory(w http.ResponseWriter, r *http.Request) {

	o, err := h.categoryC.GetAll(r.Context(), &emptypb.Empty{})
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