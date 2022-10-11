package main

import (
	"fmt"
	"os"
)
const MinimunSalary = 150000

func verifyMinimunSalary(salary int) (err error){
	if salary < MinimunSalary{
		err = fmt.Errorf("error: el mìnimo imponible es de %d y el salario ingresao es de %d", MinimunSalary, salary)
		return
	}
	fmt.Println("Debe pagar impuesto")
	return nil
}

func main(){
	salary := 100000

	err := verifyMinimunSalary(salary)

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

}