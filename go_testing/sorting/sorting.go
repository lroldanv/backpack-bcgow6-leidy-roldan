package sorting

import "sort"

func sortSlice(numbers []int) []int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	return numbers
}
