package service

import (
	"context"

	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
)

func (s *service) Create(ctx context.Context, category *domain.Category) error {
	err := s.Repo.Create(ctx, category)
	if err != nil {
		return err
	}
	return nil
}
