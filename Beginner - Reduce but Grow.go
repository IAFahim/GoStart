package main

func Grow(arr []int) int {
	var ans int=1
	for i := 0; i < len(arr); i++ {
		ans *= arr[i]
	}
	return ans
}

func main() {

}
