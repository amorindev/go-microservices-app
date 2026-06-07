package main

import (
	"context"

	"github.com/amorindev/go-microservices-app/shared"
	ordersPB "github.com/amorindev/go-microservices-app/shared/api/proto/gen"
)

type service struct {
	store OrdersStore
}

func NewSrv(store OrdersStore) *service {
	return &service{store: store}
}

func (s *service) Create(ctx context.Context) error {
	return nil
}

func (s *service) Validate(ctx context.Context,p *ordersPB.CreateOrderReq) error{
	if len(p.Items) == 0 {
		return shared.ErrNoItems
	}

	// Validate with stock service
	
	return nil
}
