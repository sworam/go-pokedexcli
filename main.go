package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var cleanedWords []string
	for _, word := range words {
		trimmed := strings.Trim(word, " ")
		if len(trimmed) > 0 {
			cleanedWords = append(cleanedWords, strings.ToLower(trimmed))
		}
	}
	return cleanedWords
}
