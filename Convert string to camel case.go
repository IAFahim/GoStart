package main

import "strings"

func ToCamelCase(s string) string {
	var sb strings.Builder
	var turn bool
	var offset = 'A' - 'a'
	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '_' {
			turn = true
			continue
		}
		if turn && 'a' <= s[i] && s[i] <= 'z' {
			sb.WriteByte(s[i] + uint8(offset))
		} else {
			sb.WriteByte(s[i])
		}
		turn = false
	}
	return sb.String()
}

func main() {
	println(ToCamelCase("The_Stealth_Warrior"))
	println(ToCamelCase("the-stealth-warrior"))
}
