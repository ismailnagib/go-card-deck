package main

import (
	"os"
	"strconv"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) < 52 {
		t.Errorf("Expected deck length of 52, instead got %v", len(d))
	}

	uniqueCard := true

	for i, card := range d {
		if !uniqueCard {
			t.Error("Expected deck content to be unique")
			break
		}

		for _, anotherCard := range d[i+1:] {
			if card == anotherCard {
				uniqueCard = false
				break
			}
		}
	}
}

func TestDeckSaveToFileAndReadFromFile(t *testing.T) {
	filename := "deckTesting.txt"
	fileSeparator := ","

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename, fileSeparator)

	savedString := readFromFile(filename)
	savedDeck := stringToDeck(savedString, fileSeparator)

	if len(savedDeck) != len(d) {
		t.Error("Expected saved deck length (" + strconv.Itoa(len(savedDeck)) + ") to be equal with original deck length (" + strconv.Itoa(len(d)) + ")")
	}

	for i, card := range d {
		if card != savedDeck[i] {
			t.Error("Expected saved deck content to be equal with original deck content")
			break
		}
	}

	os.Remove(filename)
}
