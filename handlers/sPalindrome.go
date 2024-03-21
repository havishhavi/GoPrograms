package handlers

import "fmt"

//Program to find string is palindrome or not
func IsPalindrome(data string) bool {
	revData := ""
	l := len(data)

	for i := l - 1; i >= 0; i-- {
		revData = revData + string(data[i])
	}
	fmt.Println(revData)

	if data == revData {
		return true
	} else {
		return false
	}
}
