package main

import "strings"

func toWeirdCase(str string) string {
	var sb strings.Builder
	offset := uint8('a' - 'A')
	for i, j := 0, 0; i < len(str); i++ {
		if str[i] == ' ' {
			j = 0
			sb.WriteByte(' ')
			continue
		}
		if j&1 != 0 {
			if 'A' <= str[i] && str[i] <= 'Z' {
				sb.WriteByte(str[i] + offset)
			} else {
				sb.WriteByte(str[i])
			}
		} else {
			if 'a' <= str[i] && str[i] <= 'z' {
				sb.WriteByte(str[i] - offset)
			} else {
				sb.WriteByte(str[i])
			}
		}
		j++
	}
	return sb.String()
}

func main() {
	println(toWeirdCase("This is a test Looks like you passed"))
}
