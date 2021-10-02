package main

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers); i++ {
		c := numbers[i]
		for j := 0; j < i; j++ {
			if c+numbers[j] == target {
				return [2]int{numbers[j], c}
			}
		}
		for j := i + 1; j < len(numbers); j++ {
			if c+numbers[j] == target {
				return [2]int{c, numbers[j]}
			}
		}

	}
	return [2]int{}
}

func main() {

}
