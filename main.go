package main

func main() {
	handFileName := "savedHand.txt"
	handSeparator := ","
	previousHand := readFromFile(handFileName)

	if len(previousHand) > 0 {
		println("=== PREV =========================")
		stringToDeck(previousHand, handSeparator).print()
		println("==================================")
	} else {
		println("There is not a previous hand.")
	}

	cards := newDeck()
	hand, _ := cards.deal(5)

	println()
	println("=== HAND =========================")
	hand.print()
	println("==================================")
	println()

	println(hand.saveToFile(handFileName, handSeparator))

	hand.shuffle(3)
}
