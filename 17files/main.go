package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Sohan is very smart\n"

	file, err := os.Create("./test.txt")
	checkNilError(err)

	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is", length)
	readFile("./test.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is\n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
