package main

import "strings"

func wave(words string) []string {
	at, s := make([]int, len(words)), 0
	for i := 0; i < len(words); i++ {
		if words[i] != ' ' {
			at[s] = i
			s++
		}
	}
	arr, offset := make([]string, s), uint8('a'-'A')
	for i := 0; i < s; i++ {
		n := 0
		var sb strings.Builder
		for j := 0; j < len(words); j++ {
			if words[j] != ' ' {
				if n == i {
					sb.WriteByte(words[j] - offset)
				} else {
					sb.WriteByte(words[j])
				}
				n++
			} else {
				sb.WriteByte(words[j])
			}
		}
		arr[i] = sb.String()
	}
	return arr
}

func main() {
	wave(" x yz")
}
