package port

import (
	"context"

	"github.com/amorindev/go-microservices-app/orders-srv/domain"
)

type OrdersSrv interface {
	Create(context.Context, *domain.Order) error
}

type OrdersStore interface {
	Create(context.Context, *domain.Order) error
}
