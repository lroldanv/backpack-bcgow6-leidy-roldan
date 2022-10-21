package products

import (
	"testing"
	"time"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateName(t *testing.T) {
	// Arrange

	initialDatabase := []domain.Product{
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

	expectedResult := domain.Product{
		ID:        1,
		Name:      "updatedName",
		Color:     "red",
		Price:     0,
		Stock:     0,
		Code:      "",
		Published: false,
		CreatedAt: time.Time{},
	}

	myMockStore := store.MockStorage{
		DataMock: initialDatabase,
	}
	repository := NewRepository(&myMockStore)

	// Act
	result, err := repository.UpdateName(1, "updatedName")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, myMockStore.ReadWasCalled)
}
