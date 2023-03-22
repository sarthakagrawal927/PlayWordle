package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Wordle, write PLAY to start from beginning or write the word and result to continue from last play")
	input := readWordFromCli(*reader)
	if getFormattedString(input) == "PLAY" {
		playFromStart()
	} else if strings.Contains(getFormattedString(input), "TEST_ALL") {
		simulateForAll()
	} else if strings.Contains(getFormattedString(input), "TEST") {
		isWordPlayable(input)
	} else {
		playWithCurrent(input)
	}
}

func playFromStart() {
	var status string
	var currWord string
	reader := bufio.NewReader(os.Stdin)
	currWordCount := int(0)
	copy(currentWords, wordList)
	for status != "WIN" {
		findAlphabetsFrequency()
		currWord = getMostProbableWord(currWordCount)
		fmt.Println(currWord, len(currentWords), characterCountMap)
		input := readWordFromCli(*reader)
		status = getFormattedString(input)
		moves := makePlayObject(currWord, status)
		currWordCount += 1
		playMoves(moves)
	}
}

func isWordPlayable(input string) {
	_, word := splitInput(input)

	if len(word) != 5 {
		fmt.Println("Please add 5 letter word")
		return
	}
	moves := make(chan int, 1)
	findAlphabetsFrequency()

	simulate(word, moves)
	fmt.Println(<-moves, word)
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

func playWithCurrent(input string) {
	curr, result := splitInput(input)
	findWordWithLastAnswer(curr, result)
}

func findWordWithLastAnswer(curr string, result string) {
	for idx, character := range curr {
		newPlayObj := playObject{pos: idx, char: string(character), color: string(result[idx])}
		playSingleMove(newPlayObj)
	}
	fmt.Println(getMostProbableWord(2))
}
