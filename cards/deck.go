package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

// Any variable of type "deck" now gets access to "print" func
// (d deck) is a receiver
// d -> the actual copy of the deck we're working with is available in the function as a variable called 'd' (reciever name)
// Reciever name should usually be 1 or 2 letter abbreviation to the type it refers to
// deck -> declares type that can use this function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// When u have a variable that is going to be unused, use _
	// to tell go that we don't need to use it
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Multiple return values must be declared in parentheses ()
// Inside the function return values are separated by comma
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}