package main

import ("fmt"
		"errors"
)

func calculateTaxes (salary float64) (float64, error) {

	if salary < 0 {
		return 0, errors.New("Negative value")
		err = error.
	}

	if salary > 50000 {
		return salary * 0.17
		
	}else if salary > 150_000 {
		return salary * 0.27
	}
	return 0
}

func main(){
	salary := 78000.0
	fmt.Println(calculateTaxes(salary))

}

// Generar los errores de manera global???