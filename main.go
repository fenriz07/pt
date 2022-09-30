package main

import "fmt"

func main() {
	status := validateWord("reconocer")
	fmt.Println(status)
}

func validateWord(word string) bool {

	reverseWord := ""

	for _, l := range word {
		ld := string(l)
		reverseWord = ld + reverseWord
	}

	if word == reverseWord {
		return true
	}
	return false
}
