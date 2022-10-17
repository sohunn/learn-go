package main

import (
	"fmt"
	"net/url"
)

const myURL = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=69420"

func main() {

	// parsing the url
	parsed, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}

	fmt.Println(parsed.Scheme)
	fmt.Println(parsed.Host)
	fmt.Println(parsed.Path)
	fmt.Println(parsed.RawQuery)
	fmt.Println(parsed.Port())

	params := parsed.Query()
	fmt.Println(params)

	for _, val := range params {
		fmt.Println(val)
	}

	// constructing urls from chunks

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=sohan",
	}

	fmt.Println(partsOfURL.String())
}
