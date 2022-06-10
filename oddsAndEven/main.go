package main

import "fmt"

func main() {
	//oddsAndEven := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	oddsAndEven := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range oddsAndEven {
		if n%2 == 0 {
			fmt.Printf("%d is even", n)
		} else {
			fmt.Printf("%d is odd", n)
		}
		fmt.Println("")
	}
}
