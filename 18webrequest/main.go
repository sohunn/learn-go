package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://dyno.gg/api/status"

func main() {
	resp, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Printf("Response is of type %T", resp)

	databytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(databytes))
}
