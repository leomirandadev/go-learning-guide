package main

import "fmt"

func main() {
	test := new()
	test.sayHello()
}

type exampleInterface interface {
	sayHello()
}

// This function is force the struct config implementate the functions of interface
func new() exampleInterface {
	return config{}
}

type config struct{}

// With this function, every variable typed with config, can access the function say hello
func (config) sayHello() {
	fmt.Println("Hello!")
}
