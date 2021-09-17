package main

func main() {
	/*
	cards := newDeck()

	// Since "deal" function returns 2 different values
	// We must assign them to 2 different new variables
	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()

	cards := newDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")
	cards.print()*/

	cards := newDeck()
	cards.shuffle()
	cards.print()
}