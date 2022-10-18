package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	Getdata("http://localhost:1111/get")
}

func Getdata(url string) {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var responseString strings.Builder

	data, _ := io.ReadAll(res.Body)
	byteCount, _ := responseString.Write(data)

	fmt.Println("The byte count is", byteCount)
	fmt.Println(responseString.String())
}
