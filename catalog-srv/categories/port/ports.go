package port

import (
	"context"

	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
)

type CategorySrv interface {
	Create(context.Context, *domain.Category) error
	GetAll(context.Context) ([]*domain.Category,error)
}

type CategoryRepo interface {
	Create(context.Context, *domain.Category) error
	GetAll(context.Context) ([]*domain.Category,error)
}
