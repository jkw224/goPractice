package main

import "fmt"

func main() {

	// cards := newDeck()
	// err := cards.saveToFile("my_cards")

	var cards deck

	if len(cards) == 0 {
		cards = newDeck()
		err := cards.saveToFile("my_cards")
		if err != nil {
			fmt.Println("Error --", err)
		}
	}

	fmt.Println("New deck from file: ", cards.newDeckFromFile("my_cards"))
}
