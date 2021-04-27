package main

func main() {
	cards := newDeckFromFile("myCards")
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	remainingDeck.print()
	cards.saveToFile("myCards")
	cards.print()
}
