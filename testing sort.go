package main

import (
	"fmt"
	"sort"
	"strings"
)

type employee struct {
	name   string
	salary int
}

type employeeList []employee

func (e employeeList) Len() int {
	return len(e)
}

func (e employeeList) Less(i, j int) bool {
	a := len(e[i].name)
	b := len(e[j].name)
	if a > b {
		return true
	} else if a == b {
		return 0 > strings.Compare(e[i].name, e[j].name)
	} else {
		return false
	}
}

func (e employeeList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {

	eList := make([]employee, 0, 26)
	eList = append(eList, employee{name: "yyy", salary: 3})
	eList = append(eList, employee{name: "bbbb", salary: 6})
	eList = append(eList, employee{name: "aaaa", salary: 4})
	eList = append(eList, employee{name: "mmmmmm", salary: 6})
	sort.Sort(employeeList(eList))
	for _, employee := range eList {
		fmt.Printf("Name: %s Salary %d\n", employee.name, employee.salary)
	}
}
