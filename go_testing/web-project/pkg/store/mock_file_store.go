package store

import (
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
)

type MockStorage struct {
	ReadWasCalled bool
	DataMock      []domain.Product
	ErrOnWrite    error
	ErrOnRead     error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	m.ReadWasCalled = true
	if m.ErrOnRead != nil {
		return m.ErrOnRead
	}
	// Casting data pointer --> products slice
	castedData := data.(*[]domain.Product)
	*castedData = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.ErrOnWrite != nil {
		return m.ErrOnWrite
	}
	castedData := data.(*domain.Product)
	m.DataMock = append(m.DataMock, *castedData)
	return nil
}
