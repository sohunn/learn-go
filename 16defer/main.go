package main

import "fmt"

func main() {
	defer fmt.Println("One")
	defer fmt.Println("Two")
	// LIFO: Last in first out

	fmt.Println("Hello")
	myDefer()
}

// One, two, 0, 1, 2, 3, 4

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/*
	Output:
	Hello
	4
	3
	2
	1
	0
	Two
	One

	Queue:
	One, Two, 0, 1, 2, 3, 4
*/
