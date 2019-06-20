package main

import (
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	suitLen := len(cardSuits)
	valueLen := len(cardValues)

	for len(cards) < (suitLen * valueLen) {
		rand.Seed(time.Now().UnixNano())
		suitIndex := rand.Intn(suitLen)
		valueIndex := rand.Intn(valueLen)
		newCard := cardValues[valueIndex] + " of " + cardSuits[suitIndex]
		uniqueCard := true

		for _, card := range cards {
			if card == newCard {
				uniqueCard = false
			}
		}

		if uniqueCard {
			cards = append(cards, newCard)
		}
	}

	return cards
}

func deckToString(d deck, sep string) string {
	return strings.Join([]string(d), sep)
}

func stringToDeck(s string, sep string) deck {
	if len(s) < 1 {
		return make(deck, 0)
	}

	return deck(strings.Split(s, sep))
}

func (d deck) print() {
	for i, card := range d {
		println(i+1, card)
	}
}

func (d deck) save(filename string, sep string) {
	config := getConfig()
	data := []byte(deckToString(d, sep))
	permission, error := getFilePermission(config.DeckSaveFilePermission)

	checkError("deck.save", error)

	saveToFile(filename, data, permission)
}

func deal(d deck, size int) (deck, deck) {
	validSize := size
	deckLen := len(d)

	if validSize > deckLen {
		validSize = deckLen
	}

	randSize := deckLen - validSize + 1
	rand.Seed(time.Now().UnixNano())
	start := rand.Intn(randSize)
	stop := start + validSize
	rest := append(d[:start], d[stop:]...)
	return d[start:stop], rest
}

func shuffle(d deck, round int) []deck {
	length := len(d)
	result := append(deck(nil), d...)
	results := make([]deck, round)

	for i := 0; i < round; i++ {
		previousResult := append(deck(nil), result...)

		for j := range result {
			newIndex := j

			for newIndex == j || ((result[newIndex] == previousResult[j] || result[j] == previousResult[newIndex]) && length > 3) {
				rand.Seed(time.Now().UnixNano())
				newIndex = rand.Intn(length)
			}

			result[newIndex], result[j] = result[j], result[newIndex]
		}

		results[i] = append(deck(nil), result...)
	}

	return results
}
