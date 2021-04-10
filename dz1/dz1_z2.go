package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := "Lorem ipsum dolor sit amet"
	num, occurrence := countWords(numbers)
	fmt.Println(num, occurrence) // 5 map[amet:1 dolor:1 ipsum:1 lorem:1 sit:1]
}

func countWords(sentence string) (int, map[string]int) {
	words := make(map[string]int, 5)
	sentence = strings.TrimSpace(sentence)
	sentence = strings.ReplaceAll(strings.ReplaceAll(sentence, ".", ""), ",", "")

	for _, val := range strings.Split(sentence, " ") {
		if val = strings.ToLower(val); val != "" {
			words[val]++;
		}
	}

	return len(words), words
}