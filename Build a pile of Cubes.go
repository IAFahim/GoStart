package main

func FindNb(m int) int {
	at := 0
	ans := -1
	for i := 3; m+at > 0; i++ {
		at = int((1/4.0)*float64(-(i*i*i*i)+(6*i*i*i)-(13*i*i)+(12*i))) - 1
		if m+at == 0 {
			ans = i - 2
			break
		}
	}
	if ans > -1 {
		return ans
	}
	return -1
}

func main() {
	println(FindNb(24723578342962))
}
