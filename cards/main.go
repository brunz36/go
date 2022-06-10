package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards)
	fileName := "deck.txt"

	//err := cards.saveToFile(fileName)
	//if err != nil {
	//	fmt.Println("Unable to save deck to file.")
	//}

	newCards := newDeckFromFile(fileName)
	newCards.shuffle()
	newCards.print()
}
