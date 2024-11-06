package main

import "fmt"

func main() {
	result := add(3, 5)
	fmt.Println("The result is", result)

	format, proRes, value := multiAdder(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(format, proRes, value)
}

func add(first, second int) int {
	return first + second
}

func multiAdder(values ...int) (string, int, bool) {
	total := 0
	for _, val := range values {
		total += val
	}

	return "The result is", total, true
}
