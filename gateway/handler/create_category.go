package handler

import (
	"net/http"

	"github.com/amorindev/go-microservices-app/shared"
	categoriesPB "github.com/amorindev/go-microservices-app/shared/categories/api/grpc/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) HandleCreateCategory(w http.ResponseWriter, r *http.Request) {

	var req *categoriesPB.CreateCategoryReq
	if err := shared.ReadJSON(r, &req); err != nil {
		shared.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.categoryC.Create(r.Context(), req)
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
