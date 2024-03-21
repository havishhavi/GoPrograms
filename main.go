package main

import (
	"fmt"

	"www.goprograms.com/handlers"
)

func main() {

	stringIsPalindrome := handlers.IsPalindrome("havish")

	fmt.Printf("the given string is true if palindrome: %v \n", stringIsPalindrome)

	numberIsPalindrome := handlers.CheckisPalindrome(156)
	fmt.Printf("the number is true if palindrome %v \n", numberIsPalindrome)

	s := "abcdefghij"
	fmt.Println(s[4:6])
	//convert int16 to byte
	var num int16 = 25
	num2 := byte(num)
	fmt.Printf("the num is of type %T \n", num)
	fmt.Printf("the %v and the type of number %T", num2, num2)
}
