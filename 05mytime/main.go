package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to exploring time with me!")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	someDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(someDate.Format("01-02-2006 Monday"))
}
