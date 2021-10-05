package main

var r []int

func Decode(roman string) int {
	if r == nil {
		r = make([]int, 'Z')
		r['I'] = 1
		r['V'] = 5
		r['X'] = 10
		r['L'] = 50
		r['C'] = 100
		r['D'] = 500
		r['M'] = 1000
	}
	s, c, a := 0, 0, 0
	for i := 1; i < len(roman)+1; i++ {
		c = r[roman[i-1]]
		if i < len(roman) {
			a = r[roman[i]]
		} else {
			a = 0
		}
		if c >= a {
			s += c
		} else {
			s -= c
		}
	}
	return s
}

func main() {
	println(Decode("XXI"))
}
