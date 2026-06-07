package shared

import "errors"

var (
	ErrNoItems = errors.New("items must be at least one item")
)