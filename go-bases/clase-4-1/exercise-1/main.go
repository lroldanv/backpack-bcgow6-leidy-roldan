package main

import(
	"fmt"
	"os"
)

const MinimumSalary = 150000

// Custom error 
type InvalidSalary struct{
	Message string
}

// Implement interfce Error
func (e *InvalidSalary) Error() string {
	return e.Message
}

func verifyMinimunSalary(salary int) (err error){
	if salary < MinimumSalary{
		err = &InvalidSalary {
			Message: "error: el salario ingresado no alcanza el mínimo imponible",
		}
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