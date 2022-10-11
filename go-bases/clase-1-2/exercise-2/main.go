package main

import "fmt"

func main(){
	var age int = 30
	var isEmployee bool = true
	var antiquity int = 2
	var salary float64 = 200000

	if (age > 22 && isEmployee && antiquity > 1) {
		fmt.Println("The user is suitable for a loan")
		if salary > 100000 {
			fmt.Println("Interests will be applied")
		}

	}else{
		fmt.Println("The user is not suitable for a loan")
	}


}