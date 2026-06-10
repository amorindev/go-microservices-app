package memory

import (
	"context"

	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
)

func (m *Memory) GetAll(context.Context) ([]*domain.Category, error) {
	return m.CategorySlice, nil
}
