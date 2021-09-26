package main

import "container/list"

func Anagrams(word string, words []string) []string {
	for i := 0; i < len(words); i++ {
		println(words[i])
	}
	var arr [128]int
	for i := 0; i < len(word); i++ {
		arr[word[i]]++
	}
	list := list.New()

	var brr [128]int
	for i := 0; i < len(words); i++ {
		var found int
		if len(words[i]) != len(word) {
			continue
		}
		for j := 0; j < len(arr); j++ {
			brr[j] = arr[j]
		}
		for j := 0; j < len(words[i]); j++ {
			if brr[words[i][j]] > 0 {
				brr[words[i][j]]--
				found++
			}
		}
		if found == len(word) {
			list.PushBack(words[i])
		}
	}
	var answer []string = make([]string, list.Len())
	var l = list.Len()
	if l == 0 {
		return nil
	}
	for i := 0; i < l; i++ {
		ele := list.Front()
		answer[i] = ele.Value.(string)
		list.Remove(ele)
	}
	return answer
}

func main() {
	str := Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"})
	for i := 0; i < len(str); i++ {
		println(str[i])
	}
}
