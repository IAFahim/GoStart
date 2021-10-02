package main

import "strings"

func XMasTree(height int) []string {
	arr := make([]string, height+2)
	for i := 0; i < height; i++ {
		var sb strings.Builder
		for j := 0; j < height-i-1; j++ {
			sb.WriteByte('_')
		}
		for j := 0; j < 2*i+1; j++ {
			sb.WriteByte('#')
		}
		for j := 0; j < height-i-1; j++ {
			sb.WriteByte('_')
		}
		arr[i] = sb.String()
	}
	for i := 0; i < 2; i++ {
		var sb strings.Builder
		for j := 0; j < (height - 1); j++ {
			sb.WriteByte('_')
		}
		sb.WriteByte('#')
		for j := 0; j < (height - 1); j++ {
			sb.WriteByte('_')
		}
		arr[i+height] = sb.String()
	}
	return arr
}

func main() {
	arr := XMasTree(4)
	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}
}
