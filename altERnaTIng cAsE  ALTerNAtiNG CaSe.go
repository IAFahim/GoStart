package main

import "strings"

func ToAlternatingCase(str string) string {
	var sb strings.Builder
	var offset = 'A' - 'a'
	for i := 0; i < len(str); i++ {
		var c byte = str[i]
		if 'a'<=c && c<='z' {
			sb.WriteByte(c+byte(offset))
		}else if 'A'<=c && c<='Z'{
			sb.WriteByte(c-byte(offset))
		}else {
			sb.WriteByte(c)
		}
	}
	return sb.String()
}
func main() {

}
