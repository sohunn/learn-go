package main

import "fmt"

func main() {
	result := add(3, 5)
	fmt.Println("The result is", result)

	format, proRes := multiAdder(1, 2, 3, 4, 5)
	fmt.Println(format, proRes)
}

func add(first int, second int) int {
	return first + second
}

func multiAdder(values ...int) (string, int) {
	total := 0
	for _, val := range values {
		total += val
	}

	return "The result is", total
}
