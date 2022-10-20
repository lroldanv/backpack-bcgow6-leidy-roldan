package products

import (
	"fmt"
	"time"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/store"
)

// TODO: use the Product struct in the domain and rewrite tests
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
var products []Product

type Repository interface {
	GetAll() ([]Product, error)
	Save(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error)
	Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

// repository struct implements intefaces methods
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

// Create product
func (r *repository) Save(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error) {

	var products []Product

	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	p := Product{id, name, color, price, stock, code, published, createdAt}
	products = append(products, p)
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) GetAll() ([]Product, error) {

	err := r.db.Read(&products)
	if err != nil {
		return []Product{}, err
	}
	return products, nil
}

func (r *repository) Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error) {

	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreatedAt: createdAt}

	for i, product := range products {
		if product.ID == id {
			p.ID = id
			products[i] = p
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("Product with id %d not found", id)

}

func (r *repository) Delete(id int) error {

	err := r.db.Read(&products)
	if err != nil {
		return err
	}

	var cutIndex int
	var deleted bool
	for i := range products {
		if products[i].ID == id {
			cutIndex = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("Product wit id %d does not exist", id)
	}

	products = append(products[:cutIndex], products[cutIndex+1:]...)

	if err := r.db.Write(products); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateName(id int, name string) (Product, error) {

	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}

	var product Product
	var updated bool

	for i := range products {
		if products[i].ID == id {
			products[i].Name = name
			updated = true
			product = products[i]
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("Product with id %d was NOT updated", id)
	}
	return product, nil
}
