package main

import (
	"fmt"
)

func main() {
	card := newDeck()
	// card.print()

	handCard, leftCard := deal(5, card)
	handCard.print()
	leftCard.print()

	err := handCard.saveToFile("MyDeck")
	if err != nil {
		fmt.Println(err)
	}

	newCard := newDeckFromFile("MyDeck")
	newCard.print()

	newCard.shuffle()
	newCard.print()
}
