package main

import "strconv"

func BonusTime(salary int, bonus bool) string {
	var s string
	if bonus {
		 s=strconv.Itoa(salary*10)
	}else {
		s=strconv.Itoa(salary)
	}
	return "£"+s
}

func main() {

}
