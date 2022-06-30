package main

import "fmt"

type languageBot interface {
	getGreeting() string
	getGoodbye() string
}

func printGreeting(b languageBot) {
	fmt.Println(b.getGreeting())
}

func printGoodbye(b languageBot) {
	fmt.Println(b.getGoodbye())
}
