package main

type MyString string

func (s MyString) IsUpperCase() bool {
	for i := 0; i < len(s); i++ {
		if 'Z' < s[i] {
			return false
		}
	}
	return true
}

func main() {
	s := MyString("STr")
	if MyString(s).IsUpperCase() {
		println("true")
	} else {
		println("false")

	}
}
