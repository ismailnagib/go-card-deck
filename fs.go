package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func saveToFile(filename string, data []byte, permission os.FileMode) error {
	os.Remove(filename)
	return ioutil.WriteFile(filename, data, permission)
}

func readFromFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func getFilePermission(s string) (os.FileMode, error) {
	permission, error := strconv.ParseUint(s, 0, 32)
	return os.FileMode(permission), error
}
