package main

import "fmt"

func main() {
	loginCount := 9

	var msg string

	if loginCount < 10 {
		msg = "Regular User"

	} else if loginCount > 10 {
		msg = "Watch out"

	} else {
		msg = "Somewhat okay"
	}

	fmt.Println(msg)

	if 9%2 == 0 {
		fmt.Println("8 is divisible by 2")

	} else {
		fmt.Println("The number is odd")
	}

	if num := 11; num < 10 {
		fmt.Println("Num is less than 10")

	} else {
		fmt.Println("Num is not less than 10")
	}
}
