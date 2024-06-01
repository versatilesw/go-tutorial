package main

import "fmt"

func pointerExample() {
	{
		var p *int

		fmt.Println("p = ", p)
	}
	{
		var a = 5.67
		var p = &a

		fmt.Println("Value stored in variable a = ", a)
		fmt.Println("Address of variable a = ", &a)
		fmt.Println("Value stored in variable p = ", p)
	}
	{
		var a = 100
		var p = &a

		fmt.Println("a = ", a)
		fmt.Println("p = ", p)
		fmt.Println("*p = ", *p)

		// Changing the value stored in the pointed variable through the pointer
		*p = 2000
		fmt.Println("a (after) = ", a)
	}
	{
		ptr := new(int) // Pointer to an int
		*ptr = 100

		fmt.Printf("Ptr = %#x, Ptr value = %d\n", ptr, *ptr)
	}
	{
		var a = 7.98
		var p = &a
		var pp = &p

		fmt.Println("a = ", a)
		fmt.Println("address of a = ", &a)

		fmt.Println("p = ", p)
		fmt.Println("address of p = ", &p)

		fmt.Println("pp = ", pp)

		// Dereferencing a pointer to pointer
		fmt.Println("*pp = ", *pp)
		fmt.Println("**pp = ", **pp)
	}
	{
		// var x = 67
		// var p = &x

		// var p1 = p + 1 // Compiler Error: invalid operation
	}
	{
		var a = 75
		var p1 = &a
		var p2 = &a

		if p1 == p2 {
			fmt.Println("Both pointers p1 and p2 point to the same variable.")
		}
	}
}
