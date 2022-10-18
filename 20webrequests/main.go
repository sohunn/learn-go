package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// Getdata("http://localhost:1111/get")
	PostData("http://localhost:1111/post")
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

func PostData(url string) {
	requestBody := strings.NewReader(`
		{
			"name": "Sohan",
			"surname": "Shashikumar",
			"role": "fullstack developer"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
