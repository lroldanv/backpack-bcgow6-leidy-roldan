package main

import "fmt"

func calculateAverage(values...float64) float64 {
	var result float64
	result = sum(values)/len(values)
	return result
	
}


func sum(values...float64) float64{
	var total float64
	for _, value := range values {
		total += value
	}
	return total
}

func main()  {
	fmt.Println(sum(2.0, 3.0, 4.0))
	fmt.Println()
}
// salarios y proemdios node deben ser negativos