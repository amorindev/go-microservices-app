package handler

import (
	"context"
	"log"

	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
	categoriesPB "github.com/amorindev/go-microservices-app/shared/categories/api/grpc/gen"
)

func (h *Handler) Create(ctx context.Context, req *categoriesPB.CreateCategoryReq) (*categoriesPB.Category, error) {

	log.Printf("Create category req received!!: %v", req)

	/* if err := core.ValidateItems(req); err != nil {
		return nil, err
	} */

	category := &domain.Category{
		Name: req.Name,
	}

	err := h.service.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	resp := &categoriesPB.Category{
		Id:   category.ID,
		Name: category.Name,
	}


	log.Printf("Create category res sended!!: %v", req)

	return resp, nil
}
