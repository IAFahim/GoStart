package main

import (
	"fmt"
	"strconv"
)

func PrinterError(s string) string {
	var er int= 0
	for i := 0; i < len(s); i++ {
		if s[i]>'m' {
			er++;
		}
	}
	return strconv.Itoa(er)+"/"+strconv.Itoa(len(s))
}


func main() {
	fmt.Println(PrinterError("aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz"))
	fmt.Println(PrinterError("kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu"))
}
