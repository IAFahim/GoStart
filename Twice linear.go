package main

import (
	"sort"
)

func DblLinear(n int) int {
	if a == nil {
		generate(1_000_000)
	}
	return a[n]
}

var a []int

func generate(n int) {
	n = 2*n + 1
	ar := make([]int, n)
	ar[0] = 1
	for i, x, j := 2, 0, 0; i < len(ar); i += 2 {
		x = ar[j]
		ar[i-1] = 2*x + 1
		ar[i] = 3*x + 1
		j++
	}
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		mp[ar[i]] = 1
	}
	a = make([]int, len(mp))
	i := 0
	for x, _ := range mp {
		a[i] = x
		i++
	}
	sort.Ints(a)
}

func main() {
	println(DblLinear(0))
	println(DblLinear(10))
	println(DblLinear(50))
	println(DblLinear(1000))
}
