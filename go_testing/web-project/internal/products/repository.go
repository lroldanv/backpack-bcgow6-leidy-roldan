package products

import (
	"fmt"
	"time"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/store"
)

// Global variable to store products
var products []domain.Product

type Repository interface {
	GetAll() ([]domain.Product, error)
	Save(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (domain.Product, error)
	Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
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
func (r *repository) Save(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (domain.Product, error) {

	var products []domain.Product

	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	p := domain.Product{ID: id, Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreatedAt: createdAt}
	products = append(products, p)
	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (r *repository) GetAll() ([]domain.Product, error) {

	err := r.db.Read(&products)
	if err != nil {
		return []domain.Product{}, err
	}
	return products, nil
}

func (r *repository) Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (domain.Product, error) {

	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	p := domain.Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, CreatedAt: createdAt}

	for i, product := range products {
		if product.ID == id {
			p.ID = id
			products[i] = p
			return p, nil
		}
	}
	return domain.Product{}, fmt.Errorf("product with id %d not found", id)
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
		return fmt.Errorf("product with id %d does not exist", id)
	}

	products = append(products[:cutIndex], products[cutIndex+1:]...)

	if err := r.db.Write(products); err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateName(id int, name string) (domain.Product, error) {

	err := r.db.Read(&products)
	if err != nil {
		return domain.Product{}, err
	}

	var product domain.Product
	var updated bool

	for i := range products {
		if products[i].ID == id {
			products[i].Name = name
			updated = true
			product = products[i]
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("product with id %d was NOT updated", id)
	}
	return product, nil
}
