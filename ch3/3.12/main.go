package main

import "fmt"

//Determines that the string has the same characters but in a different order
func checkStr(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)

	for _, r := range s1 {
		m[r]++
	}

	for _, r := range s2 {
		m[r]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkStr("abc", "cba"))
	fmt.Println(checkStr("abc", "cb"))
}
