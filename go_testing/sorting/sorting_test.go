package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortSlice(t *testing.T) {

	// Arrange
	slice := []int{1, 3, 2}
	expectedResult := []int{1, 2, 3}

	// Execute
	result := sortSlice(slice)

	// Validate
	assert.Equal(t, expectedResult, result, "deben ser iguales")

}
