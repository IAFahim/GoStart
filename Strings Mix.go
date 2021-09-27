package main

import (
	"sort"
	"strings"
)

type pair struct {
	name string
}

type pairList []pair

func (e pairList) Len() int {
	return len(e)
}

func (e pairList) Less(i, j int) bool {
	a := len(e[i].name)
	b := len(e[j].name)
	if a > b {
		return true
	} else if a == b {
		return 0 > strings.Compare(e[i].name, e[j].name)
	} else {
		return false
	}
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
	pairs := make([]pair, 0, 26)
	size := 0
	for i := 'a'; i <= 'z'; i++ {
		if arr[i] > 1 || brr[i] > 1 {
			var s uint8
			c := 1
			if arr[i] == brr[i] {
				s = '='
				c = arr[i]
			} else if arr[i] > brr[i] {
				s = '1'
				c = arr[i]
			} else {
				s = '2'
				c = brr[i]
			}
			var sb strings.Builder
			for j := 0; j < c; j++ {
				sb.WriteByte(byte(i))
			}
			pairs = append(pairs, pair{name: string(s) + ":" + sb.String()})
			size++
		}
	}
	sort.Sort(pairList(pairs))
	var sb strings.Builder
	for i := 0; i < len(pairs); i++ {
		sb.WriteString(string(pairs[i].name))
		if i != len(pairs)-1 {
			sb.WriteByte('/')
		}
	}
	return sb.String()
}

func main() {
	println(Mix(" In many languages", " there's a pair of functions"))
}
