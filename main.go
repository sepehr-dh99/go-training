package main

func main() {
	myDeck := newDeckFromFile("test_sepehr")

	myDeck.shuffle()
	myDeck.print()
}
