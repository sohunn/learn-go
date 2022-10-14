package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	diceNum := rand.Intn(6) + 1

	fmt.Println("The value of dice is", diceNum)

	// auto-fallthrough protection, for explicit purposes add them manually at the end

	switch diceNum {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
		fallthrough
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots and roll again")
	default:
		fmt.Println("Umm? what was that??")
	}
}
