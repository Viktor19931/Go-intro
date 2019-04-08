package main

func main() {
	// cards := newDwckFromFile("myNewFile1")
	cards := newDeck()
	cards.shuffle()
	cards.print()


	// cards.saveToFile("myNewFile")

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
}
