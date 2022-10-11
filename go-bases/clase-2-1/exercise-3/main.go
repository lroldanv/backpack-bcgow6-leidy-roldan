package main

import (
	"fmt"
	"math"
)

func calculateSalary(time float64, category string) float64 {
	var salary float64
	switch category {
	case "A":
		salary = math.Ceil(time/60) * 3000
	case "B":
		salary = math.Ceil(time/60) * 15000
	case "C":
		salary = math.Ceil(time/60) * 1000
	default:
		fmt.Println("The category does not exist")
	}
	return salary
}

func main() {
	fmt.Println(calculateSalary(60, "A"))

}
