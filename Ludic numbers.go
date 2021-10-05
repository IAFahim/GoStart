package main

var arr []int

func ludic(n int) {
	arr = make([]int, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = i
	}
	for i := 1; i < len(arr); i++ {
		first_ludic := arr[i]
		remove_index := i + first_ludic
		for remove_index < len(arr) {
			arr = append(arr[:remove_index], arr[remove_index+1:]...)
			remove_index = remove_index + first_ludic - 1
		}
	}
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
}

func SumLudic(n int) int {
	if arr == nil {
		ludic(100000)
	}
	return arr[n-1]
}

func main() {
	println(SumLudic(1))
}
