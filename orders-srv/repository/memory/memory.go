package memory

import (
	"context"
	"errors"
	"strconv"

	"github.com/amorindev/go-microservices-app/orders-srv/domain"
)

type store struct {
	OrderSlice []*domain.Order
}

func NewStore() *store {
	return &store{
		OrderSlice: []*domain.Order{},
	}
}

func (s *store) Create(ctx context.Context, order *domain.Order) error {

	if len(s.OrderSlice) == 0 {
		order.ID = "1"
	} else {
		oLength := len(s.OrderSlice)
		lastOrderID := s.OrderSlice[oLength-1].ID

		n, err := strconv.Atoi(lastOrderID)
		if err != nil {
			return errors.New("Invalid number")
		}
		order.ID = strconv.Itoa(n + 1) 
	}

	s.OrderSlice = append(s.OrderSlice, order)

	return nil
}
