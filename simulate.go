package main

import (
	"fmt"
	"strings"
)

// function to check whether the given word is winnable with the current algorithm or not, returns number of attempts if won else -1
func simulate(word string) int {
	currentWords = make([]string, len(wordList))
	copy(currentWords, wordList)
	if !wordExists(word) {
		fmt.Println("Error: word not found in the list", word)
		return -1
	}
	for currentMove := int(1); ; currentMove += 1 {
		currWord := getMostProbableWord(currentMove - 1)
		status := getComparison(word, currWord)
		findAlphabetsFrequency()
		// fmt.Println(currWord, status)
		if status == "GGGGG" {
			return currentMove
		}
		moves := makePlayObject(currWord, status)
		playMoves(moves)
	}
}

func getComparison(target string, shot string) string {
	comparison := ""
	if len(target) != len(shot) {
		fmt.Println("Error: lengths of words are not same", shot)
		return comparison
	}
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
	count := 0
	total_moves := 0
	for _, word := range wordList {
		moves := simulate(word)
		total_moves += moves
		if moves > 6 {
			count++
		}
	}
	fmt.Println("Total words:", len(wordList), "Total not winnable:", count)
	fmt.Println("Average moves:", float64(total_moves*1.0/len(wordList)))
}
