package handler

import (
	"context"
	"log"

	"github.com/amorindev/go-microservices-app/orders-srv/core"
	"github.com/amorindev/go-microservices-app/orders-srv/domain"
	ordersPB "github.com/amorindev/go-microservices-app/shared/orders/api/grpc/gen"
)

func (h *Handler) Create(ctx context.Context, req *ordersPB.CreateOrderReq) (*ordersPB.Order, error) {

	log.Printf("Create order req received!!: %v", req)

	if err := core.ValidateItems(req); err != nil {
		return nil, err
	}
	
	order := &domain.Order{
		UserID: req.UserId,
		Total:  req.Total,
	}

	items := make([]*domain.Item, 0, len(req.Items))
	for _, itemPB := range req.Items {
		item := &domain.Item{
			ProductID: itemPB.ProductId,
			Quantity:  itemPB.Quantity,
			Price:     itemPB.Price,
		}

		items = append(items, item)
	}

	order.Items = items

	err := h.service.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	resp := &ordersPB.Order{
		Id:     order.ID,
		UserId: order.UserID,
		Total:  order.Total,
	}

	itemsPB := make([]*ordersPB.Item, 0, len(order.Items))
	for _, item := range order.Items {
		itemPB := &ordersPB.Item{
			Id:        item.ID,
			ProductId: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}

		itemsPB = append(itemsPB, itemPB)
	}

	resp.Items = itemsPB

	log.Printf("Create order res sended!!: %v", req)

	return resp, nil
}