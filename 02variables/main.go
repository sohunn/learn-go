package main

import "fmt"

const PublicVariable = "hey, i am public"

func main() {
	var username string = "sohan shashikumar"
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type %T \n", isLoggedin)

	var smallval int = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type %T \n", smallval)

	var smallfloat float32 = 255.25
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type %T \n", smallfloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type %T \n", anotherVariable)

	// implicit type declaration
	var name = "sohan shashi"
	fmt.Println(name)

	// no var style with implicit type declaration
	// works only in methods and not in the global scope
	numberOfUsers := 300
	fmt.Println(numberOfUsers)
}
