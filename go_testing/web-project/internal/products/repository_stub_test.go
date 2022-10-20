package products

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s StubStore) Read(data interface{}) error {
	product1 := Product{
		ID:        1,
		Name:      "car",
		Color:     "red",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}
	product2 := Product{
		ID:        2,
		Name:      "bycicle",
		Color:     "yellow",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}
	products := []Product{product1, product2}
	MarshalData, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return json.Unmarshal(MarshalData, &data)
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	myStubStore := StubStore{}
	repository := NewRepository(myStubStore)
	product1 := Product{
		ID:        1,
		Name:      "car",
		Color:     "red",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}
	product2 := Product{
		ID:        2,
		Name:      "bycicle",
		Color:     "yellow",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}
	expectedData := []Product{product1, product2}

	// Execute
	data, err := repository.GetAll()

	assert.Equal(t, expectedData, data)
	assert.Nil(t, err)
}
