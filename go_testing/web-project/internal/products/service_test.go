package products

import (
	"errors"
	"testing"
	"time"

	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/internal/domain"
	"github.com/lroldanv/backpack-bcgow6-leidy-roldan/go_testing/web-project/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
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

	mockStorage := store.MockStorage{
		DataMock: database,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	results, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, database, results)
}

// TODO test each error individually
func TestServiceIntegrationGetAlllFail(t *testing.T) {
	// Arrange
	expectedError := errors.New("hello, I´m a error")

	mockStorage := store.MockStorage{
		DataMock:   nil,
		ErrOnWrite: nil,
		ErrOnRead:  errors.New("hello, I´m a error"),
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	results, err := service.GetAll()

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, results)
}

func TestServiceIntegrationUpdate(t *testing.T) {
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

	expectedResult := domain.Product{
		ID:        2,
		Name:      "mostCuteBycicle",
		Color:     "blue",
		Price:     10.0,
		Stock:     5,
		Code:      "b123",
		Published: true,
		CreatedAt: time.Time{},
	}

	mockStorage := store.MockStorage{
		DataMock: database,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.Update(2, "mostCuteBycicle", "blue", "b123", 10.0, 5, true, time.Time{})

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)

}

func TestServiceIntegrationDeleteOk(t *testing.T) {
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
	expectedDatabase := []domain.Product{
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
	}

	mockStorage := store.MockStorage{
		DataMock: initialDatabase,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	err := service.Delete(2)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedDatabase, mockStorage.DataMock)
}

func TestServiceIntegrationDeleteIdNotFound(t *testing.T) {
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

	mockStorage := store.MockStorage{
		DataMock: initialDatabase,
	}

	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// Act
	err := service.Delete(99)

	// Assert
	assert.NotNil(t, err)
}
