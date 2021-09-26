package main

import "fmt"

func RowSumOddNumbers(n int) int {
	var l int = n*(n-1)/2;
	var r int = l + n;
	var start int= (l*2) +1;
	var end int= (r*2) +1;
	var sum int= 0;
	for i := start; i < end; i+=2 {
		//fmt.Printf("%d ",i)
		sum+=i
	}
	return sum
}

func main() {
	fmt.Printf("%d", RowSumOddNumbers(3))
}
