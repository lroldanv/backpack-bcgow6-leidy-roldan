package division

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivisionBad(t *testing.T) {
	// Arrange
	num1 := 4.0
	num2 := 0
	expectedError := "división por cero"

	// Execute
	_, err := divide(float64(num1), float64(num2))

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expectedError)

}

func TestDivisionOk(t *testing.T) {
	// Arrange
	num1 := 4.0
	num2 := 2.0
	expectedResult := 2.0

	// Execute
	result, err := divide(float64(num1), float64(num2))

	// Assert
	assert.Equal(t, expectedResult, result, "El resultado obtenido %d, es distinto del esperado: %d", result, expectedResult)
	assert.Nil(t, err)
}
