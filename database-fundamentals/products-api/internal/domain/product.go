package domain

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" binding:"required"`
	Type  string  `json:"type" binding:"required"`
	Count uint    `json:"count" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}
