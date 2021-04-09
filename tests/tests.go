package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := calc(2, 1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result is:", result)
	}
}

func calc(a, b int) (int, error) {

	if a < b {
		return 0, errors.New("B is bigger than A")
	}

	return a - b, nil
}
