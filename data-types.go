package main

import "fmt"

func basicDataTypesExample() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)
}

func booleanExample() {
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true
}

func integerExample() {
	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	var x2 uint = 500
	var y2 uint = 4500
	fmt.Printf("Type: %T, value: %v\n", x2, x2)
	fmt.Printf("Type: %T, value: %v\n", y2, y2)
}

func floatExample() {
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	var x2 float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v", x2, x2)
}

func stringExample() {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

func dataTypesExample() {
	basicDataTypesExample()
	booleanExample()
	integerExample()
	floatExample()
	stringExample()
}
