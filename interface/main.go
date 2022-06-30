package main

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGoodbye(eb)

	printGreeting(sb)
	printGoodbye(sb)
}
