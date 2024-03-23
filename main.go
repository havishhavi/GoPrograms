package main

import (
	"fmt"

	"www.goprograms.com/handlers"
)

func main() {
	// check if given string are anagram or not
	str1 := "abxdlp"
	str2 := "pxbdal"
	isAnagram := handlers.CheckAnagram(str1, str2)
	fmt.Printf("the given two strings %v , %v are %v\n", str1, str2, isAnagram)

	//check if string is palandrome or not
	stringIsPalindrome := handlers.IsPalindrome("havish")

	fmt.Printf("the given string is true if palindrome: %v \n", stringIsPalindrome)
	//check if number is palindrome or not
	numberIsPalindrome := handlers.CheckisPalindrome(156)
	fmt.Printf("the number is true if palindrome %v \n", numberIsPalindrome)

	//search for a value in slice
	var str []rune
	str = append(str, 'h', 'a', 'v', 'i', 's', 'h')
	searchValueis := handlers.ElementSearch(str, 'h')
	fmt.Printf("the given element present in the string is %v \n", searchValueis)
	//search for the values from given number
	s := "abcdefghij"
	fmt.Println(s[4:6])

	//convert int16 to byte
	var num int16 = 25
	num2 := byte(num)
	fmt.Printf("the num is of type %T \n", num)
	fmt.Printf("the %v and the type of number %T", num2, num2)

	// b2 := 'm'
	// fmt.Printf("the num is of type %T \n", b2)

	// //FIRST INDEX IS 10 AND LAST INDEX IN 50
	// a5 := [...]int{1: 10, 5: 50}
	// fmt.Println(a5)

}
