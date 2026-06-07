package main

import (
	"context"

	ordersPB "github.com/amorindev/go-microservices-app/shared/api/proto/gen"
)

type OrdersSrv interface {
	Create(context.Context) error
	Validate(context.Context, *ordersPB.CreateOrderReq) error
}

type OrdersStore interface {
	Create(context.Context) error
}
