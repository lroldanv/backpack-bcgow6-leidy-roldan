package domain

import "time"

type Product struct {
	ID        int
	Name      string
	Color     string
	Price     float64
	Stock     uint
	Code      string
	Published bool
	CreatedAt time.Time
}
