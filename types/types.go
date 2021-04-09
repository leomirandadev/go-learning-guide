package main

import (
	"fmt" // fmt is used to print and format information
	"reflect"
)

func main() {
	basicTypes()
	mapType()
	structType()
	slicesTypes()
}

func basicTypes() {

	fmt.Println("\nbasicTypes\n")
	// ====================
	// ====== boolean ======
	// ====================

	var varBoolean bool // if you don't set the variable but inform the type, the value preset is false
	fmt.Println("varBoolean:", varBoolean, reflect.TypeOf(varBoolean))

	varBoolean2 := false // if you use ":=" to set value, the type is auto identify
	fmt.Println("varBoolean2:", varBoolean2, reflect.TypeOf(varBoolean2))

	// ====================
	// ====== string ======
	// ====================

	var varString string // if you don't set the variable but inform the type, the value preset is ""
	fmt.Println("varString:", varString, reflect.TypeOf(varString))

	varString2 := "" // if you use ":=" to set value, the type is auto identify
	fmt.Println("varString2:", varString2, reflect.TypeOf(varString2))

	// ====================
	// ====== int =========
	// ====================

	var varInt int // if you don't set the variable but inform the type, the value preset is 0
	fmt.Println("varInt:", varInt, reflect.TypeOf(varInt))

	varInt2 := 0 // if you use ":=" to set value, the type is auto identify
	fmt.Println("varInt2:", varInt2, reflect.TypeOf(varInt2))

	// ====================
	// ====== float32 =====
	// ====================

	var varFloat float32 // if you don't set the variable but inform the type, the value preset is 0
	fmt.Println("varFloat:", varFloat, reflect.TypeOf(varFloat))

	varFloat2 := 0.2 // by default, the float type auto identify is float64
	fmt.Println("varFloat2:", varFloat2, reflect.TypeOf(varFloat2))
}

func mapType() {
	fmt.Println("\nmapType\n")

	var agesNotSet map[string]int // map has a type of index and type to value

	fmt.Println(agesNotSet)

	ages := map[string]int{
		"Leonardo": 24,
		"Nicole":   24,
	}

	fmt.Println("Print all map:", ages)
	fmt.Println("Print one camp map:", "Leonardo is", ages["Leonardo"], "years old")
}

func structType() {
	fmt.Println("\nmapType\n")

	type profile struct {
		name string
		age  int
	}

	// form set 1
	var varStruct profile = profile{
		name: "Leonardo",
		age:  24,
	}
	fmt.Println(varStruct)

	// form set 2
	varStruct2 := profile{
		name: "Nicole",
		age:  24,
	}
	fmt.Println(varStruct2)
}

func slicesTypes() {
	fmt.Println("\nslicesTypes\n")

	type profile2 struct {
		name string
		age  int
	}

	var sliceInts []int = []int{0, 24, 30}
	fmt.Println(sliceInts)

	var sliceStruct []profile2 = []profile2{
		profile2{
			name: "Leonardo",
			age:  24,
		},
		profile2{
			name: "Nicole",
			age:  24,
		},
	}
	fmt.Println(sliceStruct)
}
