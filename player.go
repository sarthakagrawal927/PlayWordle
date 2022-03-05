package main

import (
	"strings"
)

func getCurrentWords() words {
	return currentWords
}

// black character - letter does not exists in our word
func handleBlack(s string) {
	var n int = 0
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
	var n int = 0
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
	var n int = 0
	for _, val := range currentWords {
		if string(val[pos]) == s {
			currentWords[n] = val
			n++
		}
	}
	currentWords = currentWords[:n]
}

func getMostProbableWord() string {
	var mpWord string
	max_power := int(0)
	power := int(0)
	for _, word := range currentWords {
		power = getWordPower(word)
		if power > max_power {
			max_power = power
			mpWord = word
		}
	}
	return mpWord
}

func play(arr []playObject) {
	for _, obj := range arr {
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
}
