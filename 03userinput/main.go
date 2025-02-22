package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "Welcome to user input"
	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating for our food!")

	// comma, ok || err, ok syntax comes in play here
	input, err := reader.ReadString('\n')

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Thanks for your rating,", input)
	}
}
