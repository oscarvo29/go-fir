package main

import "os"

func WriteToFile(filename, inp string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(inp)
	if err != nil {
		panic(err)
	}
}
