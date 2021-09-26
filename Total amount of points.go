package main

import (
	"strconv"
	"strings"
)

func Points(games []string) int {
	var score int
	for i := 0; i < len(games); i++ {
		var xy [] string = strings.Split(games[i],":")
		x, _ := strconv.ParseInt(xy[0],10,32)
		y, _ := strconv.ParseInt(xy[1],10,32)
		if x>y {
			score+=3
		}else if x==y {
			score+=1
		}
	}
	return score
}

func main() {

}
