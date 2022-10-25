package products

import (
	"fmt"
)

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

	var sellerIDExists bool
	for _, product := range prodList {
		if product.SellerID == sellerID {
			sellerIDExists = true
			break
		}
	}
	if !sellerIDExists {
		return []Product{}, fmt.Errorf("seller with id %s does not exist", sellerID)
	}
	return prodList, nil
}
