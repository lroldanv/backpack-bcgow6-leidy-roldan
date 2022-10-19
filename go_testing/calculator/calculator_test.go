package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestarOk(t *testing.T) {
	// Arrange
	num1 := 2
	num2 := 4
	expectedResult := 2

	// Execute
	result := RestarOk(num2, num1)

	// Validate

	// if result != expectedResult {
	// 	t.Errorf("Función resta() arrojo el resultado = %v, pero el esperado es %v", result, expectedResult)
	// }

	assert.Equal(t, expectedResult, result, "deben ser iguales")

}

func TestRestarBad(t *testing.T) {
	// Arrange
	num1 := 2
	num2 := 4
	expectedResult := 2

	// Execute
	result := RestarBad(num2, num1)

	// Validation

	// if result != expectedResult {
	// 	t.Errorf("Función resta() arrojo el resultado = %v, pero el esperado es %v", result, expectedResult)
	// }

	assert.Equal(t, expectedResult, result, "deben ser iguales")

}
