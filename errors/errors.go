package main

import (
	"errors"
	"fmt"
)

// Errors is use all the time on golang.
// The functions normally returns a error information
// and this is very useful when you implement unit tests
func main() {
	result, err := calc(1, 2)
	if err != nil { // nil is like a null on another languages
		fmt.Println(err)
	} else {
		fmt.Println("Result is:", result)
	}

	result2, err2 := calc(2, 1)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println("Result is:", result2)
	}

}

// (a, b int) this is the min form to say (a int, b int)
func calc(a, b int) (int, error) {

	if a < b {
		return 0, errors.New("B is bigger than A")
	}

	return a - b, nil
}
