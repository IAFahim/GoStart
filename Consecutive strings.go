package main

import "strings"

func LongestConsec(strarr []string, k int) string {
	if len(strarr) == 0 || k > len(strarr) || k <= 0 {
		return ""
	}
	arr := make([]int, len(strarr)+1)
	print("{")
	for i := 1; i < len(arr); i++ {
		arr[i] += len(strarr[i-1]) + arr[i-1]
		print("\"", strarr[i-1], "\", ")
	}
	println(k, "}")
	e := k
	m := arr[k]
	for i := k + 1; i < len(arr); i++ {
		if arr[i]-arr[i-k] > m {
			m = arr[i] - arr[i-k]
			e = i
		}
	}
	var sb strings.Builder

	for i := e - k; i < e; i++ {
		sb.WriteString(strarr[i])
	}
	return sb.String()
}

func main() {
	println(LongestConsec([]string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}, 2))
}
