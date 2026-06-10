package domain

import "time"

// Price en una sucursal tendra un presio en otra otro y la moneda
type Product struct {
	ID         string     `json:"id"`
	CategoryID string     `json:"category_id"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc"`
	Price      float64    `json:"price"`
	CreatedAt  *time.Time `json:"create_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}
