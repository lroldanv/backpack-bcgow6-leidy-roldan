package products

import (
	"testing"
	"time"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Arrange
	database := []domain.Product{
		{
			ID:        1,
			Name:      "car",
			Color:     "red",
			Price:     0,
			Stock:     0,
			Code:      "",
			Published: false,
			CreatedAt: time.Time{},
		},
		{
			ID:        2,
			Name:      "bycicle",
			Color:     "yellow",
			Price:     0,
			Stock:     0,
			Code:      "",
			Published: false,
			CreatedAt: time.Time{},
		},
	}

	myStubStore := store.StubStorage{
		DataMock: database,
	}

	repository := NewRepository(&myStubStore)

	// Act
	result, err := repository.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, database, result)

}
