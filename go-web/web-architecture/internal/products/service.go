package products

import "time"

type Service interface {
	GetAll() ([]Product, error)
	Save(name, color, code string, price float64, stock uint, published bool) (Product, error)
	Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Save(name, color, code string, price float64, stock uint, published bool) (Product, error) {

	// TODO implement a function to auto-increment the ID without calling the entire array || uuid
	products, err := s.GetAll()
	if err != nil {
		return Product{}, err
	}
	productID := len(products) + 1
	createdAt := time.Now()
	product, err := s.repository.Save(productID, name, color, code, price, stock, published, createdAt)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Update(id int, name, color, code string, price float64, stock uint, published bool, createdAt time.Time) (Product, error) {
	return s.repository.Update(id, name, color, code, price, stock, published, createdAt)
}

func (s *service) UpdateName(id int, name string) (Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) (bool, error) {
	return s.repository.Delete(id)
}
