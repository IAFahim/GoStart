package main

import "fmt"

func ExpressionMatter(a int, b int, c int) int {
	fst := a * (b + c)
	sec := a * b * c
	thr := a + (b * c)
	frt := (a + b) * c
	fiv := a + b + c
	var arr []int = []int{fst, sec, thr, frt, fiv}
	var ans int = arr[0]
	for i := 1; i < len(arr); i++ {
		if ans < arr[i] {
			ans = arr[i]
		}
	}
	return ans
}

func main() {
	fmt.Println(ExpressionMatter(1, 1, 1))
}
