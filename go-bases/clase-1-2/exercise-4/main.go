package main

import "fmt"

func main()  {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var counter int
	for key := range employees{
		if employees[key] > 21 {
			counter++
		}
	}
	fmt.Println("Number of employees older than 21:", counter)

	 // Add a new element 

	 employees["Federico"] = 25

	 // Delete an element 

	delete(employees, "Pedro")



}