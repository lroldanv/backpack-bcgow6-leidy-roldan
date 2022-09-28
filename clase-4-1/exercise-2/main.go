package main

import (
	"fmt"
	"errors"
	"os"
)
const MinimunSalary = 150000

// Declare a global variable with errors.New()
var SalaryError = errors.New("error: el salario ingresado no alcanza el mínimo imponible")

func verifyMinimunSalary(salary int) (err error){
	if salary < MinimunSalary{
		err = SalaryError
		return
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

func main(){
	salary := 200000

	err := verifyMinimunSalary(salary)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

}