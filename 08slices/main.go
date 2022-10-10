package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Pineapple"}

	fmt.Printf("The type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// used to slice from specified index
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// throws index out of bounds
	// highScores[4] = 478

	highScores = append(highScores, 69420, 7987) // works just fine cuz reallocation takes place

	sort.Ints(highScores)

	fmt.Println(highScores)

	// removing elements from slices
	var courses = []string{"react", "node js", "python", "stripe"}
	index := 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
