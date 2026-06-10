package memory

import (
	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
)

type Memory struct {
	CategorySlice []*domain.Category
}

func NewStore() *Memory {
	return &Memory{
		CategorySlice: []*domain.Category{},
	}
}
