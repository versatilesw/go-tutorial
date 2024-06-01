package utils

import "fmt"

func MapExample() {
	{
		var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
		b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

		fmt.Printf("a\t%v\n", a)
		fmt.Printf("b\t%v\n", b)
	}
	{
		var a = make(map[string]string) // The map is empty now
		a["brand"] = "Ford"
		a["model"] = "Mustang"
		a["year"] = "1964"
		// a is no longer empty
		b := make(map[string]int)
		b["Oslo"] = 1
		b["Bergen"] = 2
		b["Trondheim"] = 3
		b["Stavanger"] = 4

		fmt.Printf("a\t%v\n", a)
		fmt.Printf("b\t%v\n", b)
	}
	{
		var a = make(map[string]string)
		var b map[string]string

		fmt.Println(a == nil)
		fmt.Println(b == nil)
	}
	{
		var a = make(map[string]string)
		a["brand"] = "Ford"
		a["model"] = "Mustang"
		a["year"] = "1964"

		fmt.Println(a)

		a["year"] = "1970" // Updating an element
		a["color"] = "red" // Adding an element

		fmt.Println(a)
	}
	{
		var a = make(map[string]string)
		a["brand"] = "Ford"
		a["model"] = "Mustang"
		a["year"] = "1964"

		fmt.Println(a)

		delete(a, "year")

		fmt.Println(a)
	}
	{
		var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day": ""}

		val1, ok1 := a["brand"] // Checking for existing key and its value
		val2, ok2 := a["color"] // Checking for non-existing key and its value
		val3, ok3 := a["day"]   // Checking for existing key and its value
		_, ok4 := a["model"]    // Only checking for existing key and not its value

		fmt.Println(val1, ok1)
		fmt.Println(val2, ok2)
		fmt.Println(val3, ok3)
		fmt.Println(ok4)
	}
	{
		var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
		b := a

		fmt.Println(a)
		fmt.Println(b)

		b["year"] = "1970"
		fmt.Println("After change to b:")

		fmt.Println(a)
		fmt.Println(b)
	}
	{
		a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

		for k, v := range a {
			fmt.Printf("%v : %v, ", k, v)
		}
	}
	{
		a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

		var b []string // defining the order
		b = append(b, "one", "two", "three", "four")

		for k, v := range a { // loop with no order
			fmt.Printf("%v : %v, ", k, v)
		}

		fmt.Println()

		for _, element := range b { // loop with the defined order
			fmt.Printf("%v : %v, ", element, a[element])
		}
	}
}
