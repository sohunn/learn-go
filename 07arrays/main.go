package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Banana"
	fruitList[1] = "Apple"
	// fruitList[2] = "Jackfruit"
	fruitList[3] = "Watermelon"

	fmt.Println(fruitList)
	fmt.Println("The length of the array is", len(fruitList))
}
