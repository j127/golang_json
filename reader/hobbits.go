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

// `go run main.go < data/hobbits.json`
func Read() {
	var input []byte

	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		input = append(input, scanner.Bytes()...)
	}

	var hobbits []hobbit
	err := json.Unmarshal(input, &hobbits)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, hobbit := range hobbits {
		fmt.Printf("%s is %d years old.\n", hobbit.Name, hobbit.Age)
	}
}
