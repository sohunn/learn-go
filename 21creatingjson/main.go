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
	decodeJSON()
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

func decodeJSON() {
	jsonData := []byte(`
		{
			"coursename": "ReactJS Bootcamp",
			"Price": 229,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev","js"]
    	}
  	`)

	var lcoCourse course
	isValidJSON := json.Valid(jsonData)

	if isValidJSON {
		fmt.Println("Valid JSON")
		json.Unmarshal(jsonData, &lcoCourse)

		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Not a valid JSON")
	}

	// used for unknown values (values which cannot be predicted)
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)

	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k, v, v)
	}
}
