package product

import (
	"context"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/database-fundamentals/products-api/internal/domain"
)

type Service interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetByName(ctx context.Context, name string) (domain.Product, error) {
	product, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) Store(ctx context.Context, p domain.Product) (domain.Product, error) {
	product, err := s.repo.Store(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}
	p.ID = product.ID
	return p, nil
}
