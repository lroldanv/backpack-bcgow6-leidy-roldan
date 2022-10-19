package division

import "fmt"

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return num1 / num2, nil
}
