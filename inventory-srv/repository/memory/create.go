package memory

import (
	"context"
	"errors"
	"strconv"

	"github.com/amorindev/go-microservices-app/inventory-srv/domain"
)

func (m *Memory) Create(ctx context.Context, inventory *domain.Inventory) error {

	if len(m.InventorySlice) == 0 {
		inventory.ID = "1"
	} else {
		oLength := len(m.InventorySlice)
		lastOrderID := m.InventorySlice[oLength-1].ID

		n, err := strconv.Atoi(lastOrderID)
		if err != nil {
			return errors.New("Invalid number")
		}
		inventory.ID = strconv.Itoa(n + 1)
	}

	m.InventorySlice = append(m.InventorySlice, inventory)

	return nil
}
 