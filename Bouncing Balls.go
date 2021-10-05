package main

func BouncingBall(h, bounce, window float64) int {
	if h > 0 && h > window && bounce > 0 && 1 > bounce {
		count, at := 1, h
		for {
			at = at * bounce
			if at > window {
				count += 2
			} else {
				return count
			}
		}
	}
	return -1
}

func main() {
	println(BouncingBall(+2.000000e+000, +5.000000e-001, +1.000000e+000))
}
