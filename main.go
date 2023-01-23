package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hellflame/argparse"
)

func main() {
	// Code to play the game
	var status string
	var currWord string
	reader := bufio.NewReader(os.Stdin)
	currWordCount := int(0)
	copy(currentWords, wordList)
	for status != "WIN" {
		findAlphabetsFrequency()
		currWord = getMostProbableWord(currWordCount)
		fmt.Println(currWord, len(currentWords), characterCountMap)
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
}

func isWordPlayable() {
	parser := argparse.NewParser("wordle", "Add word", nil)
	word := parser.String("w", "word", nil)

	err := parser.Parse(nil)
	if err != nil {
		return
	}
	if len(*word) != 5 {
		fmt.Println("Please add 5 letter word")
		return
	}
	moves := make(chan int, 1)
	findAlphabetsFrequency()

	simulate(*word, moves)
	fmt.Println(<-moves, *word)
}

var showDebug bool = false

func init() {
	initWordsArray()
	findAlphabetsFrequency()
}

func initWordsArray() {
	// sort.SliceStable(wordList, func(i, j int) bool {
	// 	return getWordPower(wordList[i]) > getWordPower(wordList[j])
	// })
	currentWords = make([]string, len(wordList))
	copy(currentWords, wordList)
}
