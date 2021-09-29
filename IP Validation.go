package main

import "strconv"

func Is_valid_ip(ip string) bool {
	found := 0
	for i, pre, j := 0, -1, 0; i < len(ip); i++ {
		last := i+1 == len(ip)
		if ip[i] == '.' || last {
			if last {
				i++
			}
			s := ip[pre+1 : i]
			if 1 < len(s) && s[0] == '0' {
				return false
			}
			a, _ := strconv.ParseInt(s, 10, 32)
			if a < 0 || 255 < a {
				return false
			}
			found++
			j++
			pre = i
		} else if '0' > ip[i] || '9' < ip[i] {
			return false
		}
	}
	return found == 4
}

func main() {
	println(Is_valid_ip("123.56.89.01"))
}
