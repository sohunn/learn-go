package main

import "fmt"

func main() {

	// pointer responsible for holding integer data type
	// default value is nil
	var ptr *int
	fmt.Println(ptr)

	myNumber := 25
	// ampersand means "give me the address of myNumber and i shall store it"
	var numPointer = &myNumber

	fmt.Println("Value of numPointer is", numPointer)
	// * means "give me the value of that memory address"
	fmt.Println("Value @ memory address what numPointer holds is", *numPointer)

	*numPointer += 2
	fmt.Println("The value of myNumber is now", myNumber)

}
