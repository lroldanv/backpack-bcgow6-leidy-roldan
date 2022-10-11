package main

import (
	"errors"
	"math"
	"os"
)

type Matrix struct {
	Values      []float64
	Height      int
	Width       int
	IsQuadratic bool
}

func (matrix *Matrix) Quadratic() {
	if matrix.Height == matrix.Width {
		matrix.IsQuadratic = true
	}
	matrix.IsQuadratic = false

}

func (matrix *Matrix) set(values ...float64) error {
	if len(values) != matrix.Height*matrix.Width {
		return errors.New("The number of values do not match matrix dimensions")
		os.Exit(1)
	}
	matrix.Values = values
	return nil
}

func (matrix Matrix) calculateMaxValue() float64 {
	max := matrix.Values[0]
	for _, element := range matrix.Values {
		max = math.Max(max, element)
	}
	return max
}

func main() {

}
