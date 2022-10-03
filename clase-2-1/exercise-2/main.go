package main

import (
	"errors"
	"fmt"
	"os"
)

func calculateAverage(values ...float64) (float64, error) {
	var total float64
	for _, value := range values {
		if value < 0 {
			return 0, errors.New("error: negative values")
			os.Exit(1)
		}
		total += value
	}
	average := total / float64(len(values))
	return average, nil
}

func main() {
	average, err := calculateAverage(23, 34, 45, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(average)

}
