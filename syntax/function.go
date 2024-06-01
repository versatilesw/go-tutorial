package main

import "fmt"

func familyName(fname string) {
	fmt.Println("Hello", fname, "Refsnes")
}

func myFunction(x int, y int) int {
	return x + y
}

func myFunction2(x int, y int) (result int) {
	result = x + y
	return
}

func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func myFunction4(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func functionExample() {
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")
	fmt.Println(myFunction(1, 2))
	fmt.Println(myFunction2(1, 2))
	fmt.Println(myFunction3(5, "Hello"))
	a, b := myFunction4(5, "Hello")
	fmt.Println(a, b)
	testcount(1)
}
