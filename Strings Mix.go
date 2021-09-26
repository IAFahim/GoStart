package main

import (
	"container/list"
	"strings"
)

type pair struct {
	s   int
	str string
}

type pairList []pair

func (e pairList) Len() int {
	return len(e)
}

func (e pairList) Less(i, j int) bool {
	return e[i].str > e[j].str
}

func (e pairList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func Mix(s1, s2 string) string {

	var arr [128]int
	for i := 0; i < len(s1); i++ {
		arr[s1[i]]++
	}
	var brr [128]int
	for i := 0; i < len(s2); i++ {
		brr[s2[i]]++
	}
	al := list.New()
	size := 0
	for i := 'a'; i <= 'z'; i++ {
		if arr[i]+brr[i] > 2 {
			var s, c int
			if arr[i] == brr[i] {
				s = '='
			} else if arr[i] > brr[i] {
				s = arr[i]
			} else {
				s = brr[i]
			}
			var sb strings.Builder
			for j := 0; j < c; j++ {
				sb.WriteByte(byte(i))
			}
			al.PushBack(pair{s, sb.String()})
			size++
		}
	}
	return ""
}

func main() {
	println(Mix("mmmmm m nnnnn y&friend&Paul has heavy hats! &", "my frie n d Joh n has ma n y ma n y frie n ds n&"))
}
