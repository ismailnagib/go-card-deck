package main

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	DeckSaveFilePermission string
	HandFileName           string
	HandDeckJoinSeparator  string
	NumberOfCardsToDeal    int
	NumberOfShuffleRounds  int
	Populated              bool
}

var constant config

func readConfig() {
	file, error := ioutil.ReadFile("config.json")

	checkError("readConfig", error)

	error = json.Unmarshal(file, &constant)

	checkError("readConfig", error)

	constant.Populated = true
}

func getConfig() config {
	if !constant.Populated {
		readConfig()
	}

	return constant
}
