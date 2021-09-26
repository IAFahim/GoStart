package main

import "fmt"

func solution(str, ending string) bool {
	fmt.Println(str+" "+ending)
	if(len(str)<len(ending)){
		return false;
	}
	for i,j := len(str)-1, len(ending)-1; j >= 0 && i>=0; {
		if str[i] != ending[j]{
			return false;
		}
		i--
		j--
	}
	return true;
}

func main() {
	fmt.Println(solution("a", "z"))
}
