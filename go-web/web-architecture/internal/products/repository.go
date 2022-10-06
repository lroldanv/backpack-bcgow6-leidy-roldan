package products

import (
	"fmt"
	"time"
)

// Product struct
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Color     string    `json:"color"`
	Price     float64   `json:"price"`
	Stock     uint      `json:"stock"`
	Code      string    `json:"code"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"createdAt"`
}

// Global variable to store products
var products = []Product{
	{ID: 1, Name: "car", Color: "red", Price: 100, Stock: 4, Code: "c123", Published: true},
	{ID: 1, Name: "bicycle", Color: "red", Price: 100, Stock: 4, Code: "b123", Published: true},
}

type Repository interface {
	GetAll() ([]Product, error)
	Save(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error)
	Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) (bool, error)
}

// repository struct implements intefaces methods
type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// Create product
func (r *repository) Save(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error) {
	p := Product{id, name, color, price, stock, code, published, createdAt}
	products = append(products, p)
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreatedAt: createdAt}

	for i, product := range products {
		if product.ID == id {
			p.ID = id
			products[i] = p
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("Product %d not found", id)

}

func (r *repository) Delete(id int) (bool, error) {
	var cutIndex int
	var deleted bool
	for i := range products {
		if products[i].ID == id {
			cutIndex = i
			deleted = true
		}
	}

	if !deleted {
		return false, fmt.Errorf("Product wit id %d does not exist", id)
	}

	products = append(products[:cutIndex], products[cutIndex+1:]...)

	return true, nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {
	var product Product
	var updated bool

	for i := range products {
		if products[i].ID == id {
			products[i].Name = name
			updated = true
			product = products[i]
		}
		if !updated {
			return Product{}, fmt.Errorf("Product with id %d was NOT updated", id)
		}
		return product, nil

	}
}
