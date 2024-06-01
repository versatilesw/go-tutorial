package main

import "fmt"

func loopExamples() {
	{
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}
	{
		adj := [2]string{"big", "tasty"}
		fruits := [3]string{"apple", "orange", "banana"}
		for i := 0; i < len(adj); i++ {
			for j := 0; j < len(fruits); j++ {
				fmt.Println(adj[i], fruits[j])
			}
		}
	}
	{
		fruits := [3]string{"apple", "orange", "banana"}
		for idx, val := range fruits {
			fmt.Printf("%v\t%v\n", idx, val)
		}
	}
}
