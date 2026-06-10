package service

import (
	"time"

	"github.com/amorindev/go-microservices-app/catalog-srv/products/domain"
)

func Create(product *domain.Product) error {
	now := time.Now().UTC()
	product.CreatedAt = &now
	return nil
}
