package main

import (
	"fmt"
)

func main() {
	// var currentWords words = getCurrentWords()
	// handleBlack("a")
	// handleYellow("b", 0)
	findAlphabetsFrequency()
	fmt.Println(getMostProbableWord())
	handleBlack("A")
	handleBlack("O")
	handleBlack("S")
	handleGreen("R", 1)
	handleGreen("E", 4)
	fmt.Println(currentWords)
	handleBlack("T")
	handleGreen("I", 2)
	handleGreen("N", 3)
	fmt.Println(currentWords)
	handleBlack("U")

	fmt.Println(getMostProbableWord())
}
