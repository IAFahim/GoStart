package main

import "strings"

func ReverseWords(str string) string {
	rev := strings.Split(str, " ")
	var sb strings.Builder
	for i := 0; i < len(rev); i++ {
		for k := len(rev[i]) - 1; k >= 0; k-- {
			sb.WriteByte(rev[i][k])
		}
		if i+1 != len(rev) {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

func main() {
	print(ReverseWords("str Tst"))

}
