package port

import (
	"context"

	"github.com/amorindev/go-microservices-app/inventory-srv/domain"
)

type InventorySrv interface {
	//Find(ctx context.Context, id string) (*domain.Inventory,error)
	Create(ctx context.Context, inventory *domain.Inventory) error
}

type InventoryRepo interface {
	Create(ctx context.Context, inventory *domain.Inventory) error
}
