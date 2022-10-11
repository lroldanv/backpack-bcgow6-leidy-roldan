package main

import (
	"fmt"
	"os"
)

type Product struct {
	Id int
	Price float64
	Quantity int
}

func (product Product) printDetails() string {
	return fmt.Sprintf("%d;%.2f;%d\n", product.Id, product.Price, product.Quantity )
}

func main(){

	p1 := Product {
		Id : 1,
		Price: 234.50,
		Quantity: 23,
	}

	p2 := Product {
		Id : 2,
		Price: 676.50,
		Quantity: 2,
	}

	entry1 := p1.printDetails()
	entry2 := p2.printDetails()

	entries := entry1 + entry2
	
	
	err := os.WriteFile("../myFile.csv", []byte(entries), 0644)

	if err != nil {
		panic(err.Error)
	}
}