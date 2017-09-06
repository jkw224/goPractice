package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card: 'Ace of Diamonds', but got '%v'", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card: 'King of Spades', but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := deck.newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 card in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
