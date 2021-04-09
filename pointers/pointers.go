package main

import "fmt"

func main() {
	name := "Leonardo"
	pointerName := &name // & is used to create the pointer

	fmt.Println(name, pointerName)

	receiveCopy(name)
	fmt.Println("Print name after 'receiveCopy' function", name)

	receivePointer(pointerName)
	fmt.Println("Print name after 'receivePointer' function", name)
}

func receiveCopy(name string) {
	name = "Nicole"
}

func receivePointer(name *string) {
	*name = "Nicole" // * is used to access the information of pointer
}
