package array

import (
	"fmt"
)

func IsAnagram(s string, t string) bool {
	var letterCount1 map[rune]int
	if len(s) != len(t) {
		return false
	}
	for index, char := range s {

		fmt.Println(string(char))
		letterCount1[char] = 1

		fmt.Println(string(t[index]))
	}
	fmt.Println(letterCount1)
	return true
}
