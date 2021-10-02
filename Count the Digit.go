package main

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	count := 0
	st := strconv.Itoa(d)
	for i := 0; i <= n; i++ {
		sq := strconv.Itoa(i * i)
		count += strings.Count(sq, st)
	}
	return count
}

func main() {
	println(NbDig(10, 1))
}
