package main

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	pr(array1)
	println("----------1------------")
	pr(array2)
	println("----------2------------")
	sort.Strings(array1)
	arr := make([]string, len(array1))
	t := 0
	pre := ""
	for i := 0; i < len(array1); i++ {
		if pre == array1[i] {
			array1[i] = ""
		}
		pre = array1[i]
	}
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if len(array1[i]) > 0 && strings.Contains(array2[j], array1[i]) {
				arr[t] = array1[i]
				t++
				break
			}
		}
	}
	ans := make([]string, t)
	for i := 0; i < t; i++ {
		ans[i] = arr[i]
	}
	return ans
}

func pr(arr []string) {
	for i := 0; i < len(arr); i++ {
		print("\"" + arr[i] + "\",")
	}
}

func main() {
	arr := InArray([]string{"cod", "cod"}, []string{"cod", "cod"})
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
