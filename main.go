package main

import (
	"encoding/json"
	"fmt"

	"github.com/j127/golang_json/reader"
)

type equipment map[string]int

type Creature struct {
	Name      string    `json:"name"` // use "name" instead of "Name"
	Age       int       `json:"age"`
	Secret    string    `json:"-"` // omit this field from JSON
	Equipment equipment `json:"equipment"`
}

// `go run main.go < data/hobbits.json`
func main() {
	hobbits := []Creature{
		{"Bilbo", 111, "asdf", equipment{"ring": 1, "sword": 1}},
		{"Frodo", 33, "zxcv", equipment{"mushrooms": 5, "shoes": 0}},
	}

	// output, err := json.MarshalIndent(hobbits, "", "\t")
	_, err := json.MarshalIndent(hobbits, "", "\t")
	if err != nil {
		fmt.Println("error marshalling json", err)
		return
	}
	// fmt.Println(string(output))

	reader.Read()
}
