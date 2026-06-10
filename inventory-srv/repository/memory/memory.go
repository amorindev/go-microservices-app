package memory

import "github.com/amorindev/go-microservices-app/inventory-srv/domain"

type Memory struct {
	InventorySlice []*domain.Inventory
}

func NewRepo() *Memory {
	return &Memory{
		InventorySlice: []*domain.Inventory{},
	}
}


