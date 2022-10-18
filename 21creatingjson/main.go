package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // skip the key entirely to be not reflected in the JSON
	Tags     []string `json:"tags,omitempty"` // ignore nullable values
}

func main() {
	encodeToJSON()
}

func encodeToJSON() {

	courses := []course{
		{"ReactJS Bootcamp", 229, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "def123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 229, "LearnCodeOnline.in", "ghi123", nil},
	}

	// convert this to JSON

	json, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", json)
}
