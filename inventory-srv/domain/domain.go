package domain

type Inventory struct {
	ID            string
	VariantID     string
	Stock         int
	ReservedStock int
}
