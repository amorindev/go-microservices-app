package core

import (
	"errors"
	"fmt"

	"github.com/amorindev/go-microservices-app/shared"
	ordersPB "github.com/amorindev/go-microservices-app/shared/orders/api/grpc/gen"
)

func ValidateItems(req *ordersPB.CreateOrderReq) error {

	if req.UserId == "" {
		return errors.New("user_id is required")

	} else if req.Total <= 0 {
		return errors.New("total must be greater than 0")
	}

	if len(req.Items) == 0 {
		return shared.ErrNoItems
	}

	for i, item := range req.Items {
		if item.ProductId == "" {
			return fmt.Errorf("Item %d -productID is required", i+1)
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("Item %d - items must be have a valid quantity", i+1)
		}

		if item.Price <= 0 {
			return fmt.Errorf("Item %d - items must be have a valid price", i+1)
		}
	}
	return nil
}
