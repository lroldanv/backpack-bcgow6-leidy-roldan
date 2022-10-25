package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Unit tests

func TestServiceGetAllBySellerOk(t *testing.T) {
	// Arrange
	dataMock := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	mockRepository := MockRepository{
		DataMock: dataMock,
		Error:    nil,
	}

	service := NewService(&mockRepository)

	// Act

	products, err := service.GetAllBySeller("NeverMind")

	assert.Nil(t, err)
	assert.Equal(t, dataMock, products)

}

func TestServiceGetAllBySellerFail(t *testing.T) {
	// Arrange

	expectedError := errors.New("ups")

	mockRepository := MockRepository{
		DataMock: nil,
		Error:    errors.New("ups"),
	}

	service := NewService(&mockRepository)

	// Act
	products, err := service.GetAllBySeller("NeverMind")

	// Assert
	assert.EqualError(t, err, expectedError.Error())
	assert.Nil(t, products)

}

// Integration tests
func TestServiceIntegrationGetAllBySellerOk(t *testing.T) {
	// Arrange
	repository := NewRepository()
	service := NewService(repository)

	expectedResult := []Product{
		{ID: "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}
	// Act
	products, err := service.GetAllBySeller("FEX112AC")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, products)
}

func TestServiceIntegrationGetAllBySellerFail(t *testing.T) {
	// Arrange
	repository := NewRepository()
	service := NewService(repository)

	// Act
	_, err := service.GetAllBySeller("madeUpID")

	// Assert
	assert.NotNil(t, err)
}
