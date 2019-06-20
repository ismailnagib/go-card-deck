package main

import "os"

func checkError(funcName string, err error) {
	if err != nil {
		println(funcName, "- Error:", err.Error())
		os.Exit(1)
	}
}
