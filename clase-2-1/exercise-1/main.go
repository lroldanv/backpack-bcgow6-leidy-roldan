package main

import (
	"errors"
	"fmt"
)

var Errfoo error = errors.New("error: negative values are forbidden")

func calculateTaxes(salary float64) (float64, error) {

	if salary < 0 {
		return 0, Errfoo
	}

	if salary > 50000.0 && salary <= 150000.0 {
		fmt.Println("Taxes are applied (17%)")
		taxes := salary * 0.17
		return taxes, nil
	}

	if salary > 150000.0 {
		fmt.Println("Taxes are applied (27%)")
		taxes := salary * 0.27
		return taxes, nil
	}

	return 0, nil

}

func main() {
	salary := 90000.0

	Taxes, err := calculateTaxes(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(calculateTaxes(Taxes))

}
