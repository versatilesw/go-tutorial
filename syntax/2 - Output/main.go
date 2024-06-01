package main

import (
	"fmt"
)

func printExample() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n", j)

	fmt.Print(i, " ", j)

	var a, b = 10, 20

	fmt.Print(a, b)
}

func printlnExample() {
	var i, j string = "Hello", "World"
	fmt.Println(i, j)
}

func printfExample() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T\n", j, j)
}

func printFormatting() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	// int
	var j = 15

	fmt.Printf("%b\n", j)
	fmt.Printf("%d\n", j)
	fmt.Printf("%+d\n", j)
	fmt.Printf("%o\n", j)
	fmt.Printf("%O\n", j)
	fmt.Printf("%x\n", j)
	fmt.Printf("%X\n", j)
	fmt.Printf("%#x\n", j)
	fmt.Printf("%4d\n", j)
	fmt.Printf("%-4d\n", j)
	fmt.Printf("%04d\n", j)

	// string format
	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%16s\n", txt)
	fmt.Printf("%-16s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	// bool format
	var k = true
	var l = false

	fmt.Printf("%t\n", k)
	fmt.Printf("%t\n", l)

	// float formating
	var m = 3.141

	fmt.Printf("%e\n", m)
	fmt.Printf("%f\n", m)
	fmt.Printf("%.2f\n", m)
	fmt.Printf("%6.2f\n", m)
	fmt.Printf("%g\n", m)
}

func main() {
	printExample()
	printlnExample()
	printfExample()
	printFormatting()
}
