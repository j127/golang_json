package main

import (
	"encoding/json"
	"fmt"
)

type equipment map[string]int

type Creature struct {
	Name      string
	Age       int
	Equipment equipment
}

func main() {
	hobbits := []Creature{
		{"Bilbo", 111, equipment{"ring": 1, "sword": 1}},
		{"Frodo", 33, equipment{"mushrooms": 1, "shoes": 0}},
	}

	output, err := json.MarshalIndent(hobbits, "", "\t")
	if err != nil {
		fmt.Println("error marshalling json", err)
		return
	}
	fmt.Printf("%v", string(output))
}
