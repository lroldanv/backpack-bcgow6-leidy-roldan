package main

import (
	"fmt"
)

type Matrix struct{
	Values [][]float64
	Rows int
	Columns int
	IsQuadratic bool
	MaxValue float64
}

func (m *Matrix) set(values ... float64){
	// m.Values = make(type, 0)

	// for i := 0; i < m.Rows; i++{
	// 	for j
	// }


	if m.Rows == m.Columns {
		m.IsQuadratic = true
	}

} 

func main() {

}

