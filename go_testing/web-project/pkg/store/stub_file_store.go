package store

import (
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
)

type StubStorage struct {
	DataMock   []domain.Product
	ErrOnWrite error
	ErrOnRead  error
}

func (m *StubStorage) Read(data interface{}) (err error) {
	if m.ErrOnRead != nil {
		return m.ErrOnRead
	}
	// Casting data pointer --> products slice
	castedData := data.(*[]domain.Product)
	*castedData = m.DataMock
	return nil
}

func (m *StubStorage) Write(data interface{}) (err error) {
	if m.ErrOnWrite != nil {
		return m.ErrOnWrite
	}
	castedData := data.(*domain.Product)
	m.DataMock = append(m.DataMock, *castedData)
	return nil
}
