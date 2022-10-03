package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const (
	minimun = "minimun"
	average = "average"
	maximun = "maximun"
)

var errFoo = errors.New("error: negative values are forbidden")

func calculateMinimun(values ...float64) (float64, error) {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < 0 {
			return 0, errFoo
		}
		min = math.Min(min, values[i])
	}
	return min, nil
}

func calculateMaximun(values ...float64) (float64, error) {
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < 0 {
			return 0, errFoo
		}
		min = math.Max(min, values[i])
	}
	return min, nil
}

func calculateAverage(values ...float64) (float64, error) {
	var total float64
	for _, value := range values {
		if value < 0 {
			return 0, errFoo
			os.Exit(1)
		}
		total += value
	}
	average := total / float64(len(values))
	return average, nil
}

func getFunction(option string) (func(...float64) (float64, error), error) {
	switch option {
	case minimun:
		return calculateMinimun, nil
	case maximun:
		return calculateMaximun, nil
	case average:
		return calculateAverage, nil
	default:
		return nil, errors.New("The function has not been defined")
	}
}

func main() {

	minFunc, err := getFunction(minimun)
	if err != nil {
		fmt.Println(err)
		return
	}
	minValue, err := minFunc(2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The minimun values is ", minValue)

	maxFunc, err := getFunction(maximun)
	if err != nil {
		fmt.Println(err)
		return
	}
	maxValue, err := maxFunc(3, 7, 8, 9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The minimun values is ", maxValue)

	averageFunc, err := getFunction(average)
	if err != nil {
		fmt.Println(err)
	}
	averageValue, err := averageFunc(2, 4, 7, 8, 9)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The minimun values is ", averageValue)

}
