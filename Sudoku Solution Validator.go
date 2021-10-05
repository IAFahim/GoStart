package main

func ValidateSolution(m [][]int) bool {
	for y := 0; y < len(m); y++ {
		sum := 0
		for x := 0; x < len(m[y]); x++ {
			sum += m[y][x]
		}
		if sum != 45 {
			return false
		}
	}

	for i := 0; i < 3; i++ {
		sum := 0
		for y := i * 3; y < (3)*(i+1); y++ {
			for x := i * 3; x < 3*(i+1); x++ {
				sum += m[y][x]
			}
		}
		if sum != 45 {
			return false
		}
	}
	return true
}

func main() {
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	println(ValidateSolution(testTrue))
}
