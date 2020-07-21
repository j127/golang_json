package reader

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type hobbit struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Read() {
	var input []byte

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		input = append(input, scanner.Bytes()...)
	}

	var hobbits []hobbit

	if err := json.Unmarshal(input, &hobbits); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Data from a JSON file\n---------------------")
	for _, hobbit := range hobbits {
		fmt.Printf("%s is %d years old.\n", hobbit.Name, hobbit.Age)
	}
}
