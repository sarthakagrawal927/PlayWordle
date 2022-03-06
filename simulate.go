package main

import (
	"fmt"
	"strings"
)

// function to check whether the given word is winnable with the current algorithm or not, returns number of attempts if won else -1
func simulate(word string) int {
	for currentMove := int(1); currentMove <= 6; currentMove += 1 {
		currWord := getMostProbableWord(currentMove - 1)
		status := getComparison(word, currWord)
		fmt.Println(currWord, status)
		if status == "GGGGG" {
			return currentMove
		}
		moves := makePlayObject(currWord, status)
		playMoves(moves)
	}
	return -1
}

func getComparison(target string, shot string) string {
	comparison := ""
	for i := 0; i < len(target); i++ {
		if target[i] == shot[i] {
			comparison += "G"
		} else if strings.Contains(target, string(shot[i])) {
			comparison += "Y"
		} else {
			comparison += "B"
		}
	}
	return comparison
}

func simulateForAll() {
	allWords := wordList
	for _, word := range wordList {
		moves := simulate(word)
		currentWords = allWords
		fmt.Println(word, moves, len(wordList))
	}
}
