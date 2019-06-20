package main

import (
	"io/ioutil"
	"math/rand"
	"os"
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

func readFromFile(filename string) string {
	content, error := ioutil.ReadFile(filename)

	if error != nil {
		println("Error: " + error.Error())
		os.Exit(1)
	}

	data := string(content)
	return data
}

func stringToDeck(str string, sep string) deck {
	arrOfStr := strings.Split(str, sep)
	return deck(arrOfStr)
}

func (d deck) print() {
	for i, card := range d {
		println(i+1, card)
	}
}

func (d deck) deal(size int) (deck, deck) {
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

func (d deck) saveToFile(filename string, sep string) string {
	arrOfStr := []string(d)
	str := strings.Join(arrOfStr, sep)
	arrOfByte := []byte(str)
	error := ioutil.WriteFile(filename, arrOfByte, 0666)

	if error != nil {
		return "Error: " + error.Error()
	}

	return "The hand was successfully saved."
}

func (d deck) shuffle(round int) deck {
	length := len(d)
	result := append(deck(nil), d...)

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

		println()
		println("=== HAND ========================= SHUFFLE", i+1)
		result.print()
		println("==================================")
	}

	return result
}
