package handlers

import (
	"fmt"
	"sort"
)

func CheckAnagram(str1, str2 string) bool {

	s1 := []rune(str1)
	s2 := []rune(str2)
	//i := 0
	if len(str1) != len(str2) {
		return false
	}

	fmt.Println(s1, ",", s2)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})

	for k := range s1 {
		if s1[k] != s2[k] {
			return false
		}

	}
	return true

}
