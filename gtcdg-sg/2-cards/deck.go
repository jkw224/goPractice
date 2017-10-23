package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

// Prints cards in deck
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Creates a new deck of 52 cards
func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Deals out x number of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Shuffle Cards
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}

// Converts deck to string
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// Save cards to file
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// Create deck from file
// returns []byte & error
func (d deck) newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error -- ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}
