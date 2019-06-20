package main

func main() {
	config := getConfig()
	data, _ := readFromFile(config.HandFileName)

	if len(config.HandDeckJoinSeparator) < 1 {
		config.HandDeckJoinSeparator = ","
	}

	previousHand := stringToDeck(string(data), config.HandDeckJoinSeparator)

	if len(previousHand) > 0 {
		println("=== PREV =========================")
		previousHand.print()
		println("==================================")
	} else {
		println("There is not a previous hand.")
	}

	cards := newDeck()
	hand, _ := deal(cards, config.NumberOfCardsToDeal)

	println()
	println("=== HAND =========================")
	hand.print()
	println("==================================")

	println()
	if len(config.HandFileName) < 1 {
		println("The hand was not saved because no filename was set on the configuration file (config.json).")
	} else {
		hand.save(config.HandFileName, config.HandDeckJoinSeparator)
		println("The hand was successfully saved.")
	}

	shuffled := shuffle(hand, config.NumberOfShuffleRounds)
	for i, hand := range shuffled {
		println()
		println("=== SHUFFLE ROUND", i+1, "==============")
		hand.print()
		println("==================================")
	}
}
