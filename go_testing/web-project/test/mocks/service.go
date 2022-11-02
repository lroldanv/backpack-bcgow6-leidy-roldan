package mocks

import (
	"fmt"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
)

type MockService struct {
	DataMock []domain.Product
	Error    string
}

func (m *MockService) GetAll() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil

}

func (m *MockService) Store(name, color, code string, price float64, stock uint, published bool) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}
	p := domain.Product{
		Name:      name,
		Color:     color,
		Code:      code,
		Price:     price,
		Stock:     stock,
		Published: published,
	}
	return p, nil
}
