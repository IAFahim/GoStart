package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(sc, &s)
	println(s)
}
