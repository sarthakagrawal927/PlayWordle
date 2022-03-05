package main

import (
	"fmt"
)

func main() {
	// var currentWords words = getCurrentWords()
	// handleBlack("a")
	// handleYellow("b", 0)
	fmt.Println(getMostProbableWord())
	handleGreen("c", 1)
	fmt.Println(getMostProbableWord())
	findAlphabetsFrequency()
	sortWordsByPossibility()
	fmt.Println(currentWords, characterCountMap)
}
