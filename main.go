package main

import (
	"fmt"
)

func main() {
	findAlphabetsFrequency()
	fmt.Println(getMostProbableWord())
	moves := []playObject{ //pos,char,color
		{0, "A", "B"},
		{1, "R", "B"},
		{2, "O", "B"},
		{3, "S", "Y"},
		{4, "E", "G"},
	}
	play(moves)
	fmt.Println(getMostProbableWord())
}
