package main

import "fmt"

func main() {
	functionNotExported()
	FunctionExported()

	functionAttr("leo")

	returnOne := functionOneReturn("leo") // if you use ":=" the type auto set is the return of function
	fmt.Println(returnOne)

	msgName, msgAge := functionTwoOrMoreReturn("leo", "24") // if you use ":=" the type auto set is the return of function
	fmt.Println(msgName, msgAge)

}

func functionNotExported() {} // if you use de first letter on lowercase, the function can be use only in this package

func FunctionExported() {} // if you use de first letter on uppercase, the function can be use without this package

func functionAttr(name string) {
	fmt.Println(name)
}

func functionOneReturn(name string) string {
	return "The name is " + name
}

func functionTwoOrMoreReturn(name string, age string) (string, string) {
	return "The name is " + name, "The age is" + age
}
