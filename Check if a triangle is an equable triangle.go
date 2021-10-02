package main

import "math"

func EquableTriangle(a, b, c int) bool {
	s := float64(a+b+c) * 0.5
	ar := math.Sqrt(s * (s - float64(a)) * (s - float64(b)) * (s - float64(c)))
	return s*2.0 == ar
}

func main() {
	EquableTriangle(5, 12, 13)
}
