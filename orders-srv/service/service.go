package service

import (
	"context"

	"github.com/amorindev/go-microservices-app/orders-srv/domain"
	"github.com/amorindev/go-microservices-app/orders-srv/port"
)

type service struct {
	store port.OrdersStore
}

func NewSrv(store port.OrdersStore) *service {
	return &service{store: store}
}

func (s *service) Create(ctx context.Context, order *domain.Order) error {
	
	err := s.store.Create(ctx, order)
	if err != nil {
		return err
	}
	
	return nil
}

/* func (s *service) Validate(ctx context.Context,p *ordersPB.CreateOrderReq) error{
	if len(p.Items) == 0 {
		return shared.ErrNoItems
	}

	// Validate with stock service, ver

	s.store.Create(ctx,)

	return nil
} */
