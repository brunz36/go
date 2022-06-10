package main

import "fmt"

func main() {
	var bar map[string]string
	fmt.Println(bar)

	foo := make(map[string]string)
	foo["red"] = "#ff0000"
	fmt.Println(foo)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
		"black": "#000000",
	}

	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Color [%s] Hex [%s] \n", color, hex)
	}
}
