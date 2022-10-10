package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println(languages)
	fmt.Println("JS is the shortform for", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	// looping over keys and values of a map
	for key, val := range languages {
		fmt.Printf("For key %v, the value is %v\n", key, val)
	}
}
