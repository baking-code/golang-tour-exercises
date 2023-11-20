package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	// split string into words (whitespace)
	split := strings.Fields(s)
	for _, word := range split {
		// accesses map
		_, ok := wordMap[word]
		if ok {
			// if exists, increment
			wordMap[word] = wordMap[word] + 1
		} else {
			// else intiialise count
			wordMap[word] = 1
		}
	}
	return wordMap
}

func main() {
	// passes function into the tour package's wordcount test suite
	wc.Test(WordCount)
}
