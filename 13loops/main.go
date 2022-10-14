package main

import "fmt"

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// traditional for loop
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// type 2
	for i := range days {
		fmt.Println(days[i])
	}

	// type 3
	for index, value := range days {
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

	rogueValue := 1
	for rogueValue <= 10 {

		if rogueValue == 2 {
			goto lco // similar to break
		}

		// never comes here
		if rogueValue == 5 {
			break
		}

		fmt.Println("The value is:", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at lco.in")
}
