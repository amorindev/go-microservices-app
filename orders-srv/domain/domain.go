package domain

type Order struct {
	ID     string
	UserID string
	Items  []*Item
	Total  float64
}

type Item struct {
	ID        string
	ProductID string
	Quantity  int32
	Price     float64
}
