package main

import "strings"

func DNAtoRNA(dna string) string {
	var str strings.Builder
	for i := 0; i < len(dna); i++ {
		if dna[i] == 'T' {
			str.WriteByte('U')
		}else {
			str.WriteByte(dna[i])
		}
	}
	return str.String()
}

func main() {

}
