package main

import "sort"

func SortNumbers(numbers []int) []int {
	if numbers == nil || len(numbers) == 0 {
		return []int{}
	}
	sort.Ints(numbers)
	return numbers
}

func main() {

}
