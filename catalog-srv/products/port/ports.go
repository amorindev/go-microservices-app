package port

import (
	"context"

	"github.com/amorindev/go-microservices-app/catalog-srv/products/domain"
)

type ProductSrv interface {
	Find(ctx context.Context, id string)
	Create(ctx context.Context, product domain.Product) error
}

type InventoryRepo interface {
	Create(ctx context.Context, product domain.Product) error
}