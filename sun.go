package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("testcase.txt")
	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	var s, sum int
	for true {
		_, err = fmt.Fscan(file, &s)
		if err == io.EOF {
			break
		}
		sum += s
	}

	println(sum)

}
