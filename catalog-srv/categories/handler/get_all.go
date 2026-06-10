package handler

import (
	"context"

	categoriesPB "github.com/amorindev/go-microservices-app/shared/categories/api/grpc/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) GetAll(ctx context.Context, _ *emptypb.Empty) (*categoriesPB.GetAllResp, error) {
	categories, err := h.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	categoriesPBSlice := make([]*categoriesPB.Category, 0, len(categories))
	for _, c := range categories {
		categoryPB := &categoriesPB.Category{
			Id:   c.ID,
			Name: c.Name,
		}

		categoriesPBSlice = append(categoriesPBSlice, categoryPB)
	}

	resp := &categoriesPB.GetAllResp{
		Products: categoriesPBSlice,
	}

	return resp, nil
}

