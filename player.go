package main

import (
	"fmt"
	"strings"
)

// black character - letter does not exists in our word
func handleBlack(s string) {
	n := 0
	for _, val := range currentWords {
		if !strings.Contains(string(val), s) {
			currentWords[n] = val
			n++
		}
	}
	currentWords = currentWords[:n]
}

// yellow character - letter exists, but not at the current position
func handleYellow(s string, pos int) {
	n := 0
	for _, val := range currentWords {
		if strings.Contains(string(val), s) && string(val[pos]) != s {
			currentWords[n] = val
			n++
		}
	}
	currentWords = currentWords[:n]
}

// green character - letter exists, at current position (or more)
func handleGreen(s string, pos int) {
	n := 0
	for _, val := range currentWords {
		if string(val[pos]) == s {
			currentWords[n] = val
			n++
		}
	}
	currentWords = currentWords[:n]
}

// main logic
func getMostProbableWord(currWordCount int) string {
	// complete alternate of the first word
	if currWordCount == 1 {
		return "UNLIT"
	}
	mpWord := ""
	max_power := float32(0)
	power := float32(0)
	for _, word := range currentWords {
		power = getWordPower(word)
		if power > max_power {
			max_power = power
			mpWord = word
		}
	}
	if showDebug {
		fmt.Println(characterCountMap, mpWord, max_power, currWordCount)
	}
	return mpWord
}

func playMoves(arr []playObject) {
	for _, obj := range arr {
		playSingleMove(obj)
	}
}

func playSingleMove(obj playObject) {
	switch obj.color {
	case "Y":
		handleYellow(obj.char, obj.pos)
	case "B":
		handleBlack(obj.char)
	case "G":
		handleGreen(obj.char, obj.pos)
	default:
	}
}

func makePlayObject(word string, colors string) []playObject {
	moves := []playObject{}
	if colors == "WIN" {
		return moves
	}
	for idx := range word {
		moves = append(moves, playObject{idx, string(word[idx]), string(colors[idx])})
	}
	return moves
}
