package service

import (
	"context"

	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
)

func (s *service) GetAll(ctx context.Context) ([]*domain.Category, error) {
	return s.Repo.GetAll(ctx)
}
