package memory

import (
	"context"
	"errors"
	"strconv"

	"github.com/amorindev/go-microservices-app/catalog-srv/categories/domain"
)

func (m *Memory) Create(ctx context.Context, category *domain.Category) error {

	if len(m.CategorySlice) == 0 {
		category.ID = "1"
	} else {
		oLength := len(m.CategorySlice)
		lastOrderID := m.CategorySlice[oLength-1].ID

		n, err := strconv.Atoi(lastOrderID)
		if err != nil {
			return errors.New("Invalid number")
		}
		category.ID = strconv.Itoa(n + 1)
	}

	m.CategorySlice = append(m.CategorySlice, category)

	return nil
}