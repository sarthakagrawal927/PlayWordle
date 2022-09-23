package main

import (
	"fmt"
	"strings"
	"time"
)

// function to check whether the given word is winnable with the current algorithm or not, returns number of attempts if won else -1
func simulate(word string, result chan int) {
	currentWords = make([]string, len(wordList))
	copy(currentWords, wordList)
	if !wordExists(word) {
		fmt.Println("Error: word not found in the list", word)
		result <- -1
		return
	}
	for currentMove := 0; ; currentMove += 1 {
		// findAlphabetsFrequency()
		currWord := getMostProbableWord(currentMove)
		statusChannel := make(chan string, 1)
		go getComparison(word, currWord, statusChannel)
		status := <-statusChannel
		if status == "GGGGG" {
			result <- currentMove
			return
		}
		moves := makePlayObject(currWord, status)
		playMoves(moves)
	}
}

func getComparison(target string, shot string, comparisonChannel chan string) {
	comparison := []byte{'B', 'B', 'B', 'B', 'B'}
	if len(target) != len(shot) {
		comparisonChannel <- "GGGGG"
		return
	}
	for i := 0; i < len(target); i++ {
		if target[i] == shot[i] {
			comparison[i] = 'G'
		} else if strings.Contains(target, string(shot[i])) {
			comparison[i] = 'Y'
		}
	}
	comparisonChannel <- string(comparison)
}

func simulateForAll() {
	showDebug = false
	startTime := time.Now()
	lostCount := 0
	win_moves := 0.0
	loss_moves := 0.0
	for _, word := range wordList {
		movesChannel := make(chan int, 1)
		go simulate(word, movesChannel)
		moves := <-movesChannel
		if moves > 6 {
			lostCount++
			fmt.Println(word, moves)
			loss_moves += float64(moves)
		} else {
			win_moves += float64(moves)
		}
	}
	elapsedTime := time.Since(startTime)
	fmt.Println("Success Percentage", float64(len(wordList)-lostCount)/float64(len(wordList))*100)
	fmt.Println("Average Win moves:", win_moves/float64(len(wordList)-lostCount))
	fmt.Println("Average Loss moves:", loss_moves/float64(lostCount))
	fmt.Println("Time taken per word", float64(elapsedTime)/float64(len(wordList)*1000000), "ms")
}
