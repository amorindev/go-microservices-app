package memory

import (
	"context"
	"errors"
	"strconv"

	"github.com/amorindev/go-microservices-app/catalog-srv/products/domain"
)

type store struct {
	ProductSlice []*domain.Product
}

func NewStore() *store {
	return &store{
		ProductSlice: []*domain.Product{},
	}
}

func (s *store) Create(ctx context.Context, product *domain.Product) error {

	if len(s.ProductSlice) == 0 {
		product.ID = "1"
	} else {
		oLength := len(s.ProductSlice)
		lastOrderID := s.ProductSlice[oLength-1].ID

		n, err := strconv.Atoi(lastOrderID)
		if err != nil {
			return errors.New("Invalid number")
		}
		product.ID = strconv.Itoa(n + 1)
	}

	s.ProductSlice = append(s.ProductSlice, product)

	return nil
}