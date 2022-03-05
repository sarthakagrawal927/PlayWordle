package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	findAlphabetsFrequency()

	var status string
	var currWord string
	reader := bufio.NewReader(os.Stdin)

	for status != "WIN" {
		currWord = getMostProbableWord()
		fmt.Println(currWord, len(currentWords))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again", err)
			return
		}
		status = strings.ToUpper(strings.TrimSuffix(input, "\n"))
		moves := makePlayObject(currWord, status)
		playMoves(moves)
	}

}
