package main

func main() {

	cards := newDeckFromFile("cards.txt")
	//cards.print()

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	cards.shuffle()
	cards.print()
	//cards.saveToFile("cards.txt")

}

func newCard() string {
	return "Five of Diamonds"
}
