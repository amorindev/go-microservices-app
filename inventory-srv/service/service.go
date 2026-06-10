package service

import (
	"context"

	"github.com/amorindev/go-microservices-app/inventory-srv/domain"
	"github.com/amorindev/go-microservices-app/inventory-srv/port"
)

type service struct {
	Repo port.InventoryRepo
}

func NewSrv(repo port.InventoryRepo) *service {
	return &service{Repo: repo}
}

func (s *service) Create(ctx context.Context, inventory *domain.Inventory) error {
	return nil
}


