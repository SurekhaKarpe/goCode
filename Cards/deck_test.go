package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	deck := newDeck()

	if len(deck) != 52 {
		t.Errorf("Expected deck length is 52, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spade" {
		t.Errorf("Expected card is 'Ace of Spade', but got %v", deck[0])
	}

	if deck[len(deck)-1] != "King of Club" {
		t.Errorf("Expected card is 'King of Club', but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_deckTesting")

	deck := newDeck()
	err := deck.saveToFile("_deckTesting")
	if err != nil {
		t.Errorf("Failed to Save file!!")
	}

	newDeck := newDeckFromFile("_deckTesting")
	if len(newDeck) != 52 {
		t.Errorf("Expected 52 card, but got %v", len(newDeck))
	}

	os.Remove("_deckTesting")
}
