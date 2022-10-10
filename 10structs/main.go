package main

import "fmt"

func main() {

	sohan := User{"sohan", "sohanshashi0884@gmail.com", true, 19}
	fmt.Println(sohan)
	fmt.Printf("Sohan's details are %+v\n", sohan)

	fmt.Printf("Sohan's age is %v\n", sohan.Age)
}

type User struct {
	Name     string
	Email    string
	Verified bool
	Age      int
}
