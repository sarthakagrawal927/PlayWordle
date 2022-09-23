package main

var characterCountMap map[rune]float32

func findAlphabetsFrequency() {
	characterCountMap = make(map[rune]float32)
	for _, val := range currentWords {
		addWordToFreq(val)
	}
}

func addWordToFreq(s string) {
	// if a word contains the same character twice, I am counting it only once
	mapCheck := make(map[rune]bool)
	for _, c := range s {
		if !mapCheck[c] {
			characterCountMap[c] += 1
			mapCheck[c] = true
		}
	}
}

func getWordPower(s string) float32 {
	power := float32(0)
	mapCheck := make(map[rune]bool)
	for _, c := range s {
		if !mapCheck[c] {
			power += characterCountMap[c]
			mapCheck[c] = true
		}
	}
	return power
}

// checks whether the word exists or not in our wordList
func wordExists(s string) bool {
	for _, val := range wordList {
		if s == val {
			return true
		}
	}
	return false
}
