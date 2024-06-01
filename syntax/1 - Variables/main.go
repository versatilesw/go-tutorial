package main

import "fmt"

var globalA int
var globalB int = 2
var globalC = 3

// Since := is used outside of a function, running the program results in an error
// a := 1

// Declaring a Constant
const PI = 3.14
const TPYED_PI float64 = 3.14

/* The code below will print Hello World */

func main() {
	// This is a comment
	fmt.Println("Hello World!")

	// Variable Declaration With Initial Value
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Value Assignment After Declaration
	var student3 string
	student3 = "John"
	fmt.Println(student3)

	globalA = 1
	fmt.Println(globalA)
	fmt.Println(globalB)
	fmt.Println(globalC)

	// Go Multiple Variable Declaration
	// Note: If you use the type keyword, it is only possible to declare one type of variable per line.
	var d, e, f, g int = 1, 3, 5, 7

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// If the type keyword is not specified, you can declare different types of variables in the same line:
	var h, i = 6, "Hello"
	j, k := 7, "World!"

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	// Go Variable Declaration in a Block
	var (
		l int
		m int    = 1
		n string = "hello"
	)

	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)

	// Go Variable Naming Rules
	// Camel Case
	var myVariableName = "John"
	fmt.Println(myVariableName)

	// Pascal Case
	var MyVariableName = "John"
	fmt.Println(MyVariableName)

	// Snake Case
	var my_variable_name = "John"
	fmt.Println(my_variable_name)

	fmt.Println(PI)
}
