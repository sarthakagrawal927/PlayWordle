package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	findAlphabetsFrequency()

	// Code to play the game
	var status string
	var currWord string
	reader := bufio.NewReader(os.Stdin)
	currWordCount := int(0)
	for status != "WIN" {
		currWord = getMostProbableWord(currWordCount)
		fmt.Println(currWord, len(currentWords))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			return
		}
		status = strings.ToUpper(strings.TrimSuffix(input, "\n"))
		moves := makePlayObject(currWord, status)
		currWordCount += 1
		playMoves(moves)
	}

	// Code to simulate the game
	// simulateForAll()
}
