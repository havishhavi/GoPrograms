package handlers

import "fmt"

//check if giver number is palindrome

func CheckisPalindrome(num int) bool {
	revNum := 0
	org := num
	for num > 0 {
		reminder := num % 10
		revNum = revNum*10 + int(reminder)
		num = num / 10

		fmt.Printf("the reminder everytime is %v \n and revNumber is %v \n", reminder, revNum)
		fmt.Printf("number in each iteration %v \n", num)

	}
	fmt.Println(revNum)
	fmt.Println(org)
	return org == revNum

}
