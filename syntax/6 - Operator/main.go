package main

import "fmt"

func main() {
	{
		var (
			sum1 = 100 + 50    // 150 (100 + 50)
			sum2 = sum1 + 250  // 400 (150 + 250)
			sum3 = sum2 + sum2 // 800 (400 + 400)
		)
		fmt.Println(sum3)
	}
	{
		var x = 10
		x += 5
		fmt.Println(x)
	}
	{
		var x = 5
		var y = 3
		fmt.Println(x > y) // returns 1 (true) because 5 is greater than 3
	}
	{
		if 20 > 18 {
			fmt.Println("20 is greater than 18")
		}
	}
	{
		x := 20
		y := 18
		if x > y {
			fmt.Println("x is greater than y")
		}
	}
	{
		time := 20
		if time < 18 {
			fmt.Println("Good day.")
		} else {
			fmt.Println("Good evening.")
		}
	}
	{
		time := 22
		if time < 10 {
			fmt.Println("Good morning.")
		} else if time < 20 {
			fmt.Println("Good day.")
		} else {
			fmt.Println("Good evening.")
		}
	}
	{
		num := 20
		if num >= 10 {
			fmt.Println("Num is more than 10.")
			if num > 15 {
				fmt.Println("Num is also more than 15.")
			}
		} else {
			fmt.Println("Num is less than 10.")
		}
	}
	{
		day := 4

		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("Not a weekday")
		}
	}
	{
		// a := 3
		// switch a {
		// case 1:
		// 	fmt.Println("a is one")
		// case "b":
		// 	fmt.Println("a is b")
		// }
	}
	{
		day := 5

		switch day {
		case 1, 3, 5:
			fmt.Println("Odd weekday")
		case 2, 4:
			fmt.Println("Even weekday")
		case 6, 7:
			fmt.Println("Weekend")
		default:
			fmt.Println("Invalid day of day number")
		}
	}
}
