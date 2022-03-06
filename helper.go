package main

var characterCountMap map[string]int

func findAlphabetsFrequency() {
	characterCountMap = make(map[string]int)
	for _, val := range currentWords {
		addWordToFreq(val)
	}
}

func addWordToFreq(s string) {
	// if a word contains the same character twice, I am counting it only once
	mapCheck := make(map[string]bool)
	for _, c := range s {
		if mapCheck[string(c)] != true {
			if characterCountMap[string(c)] == 0 {
				characterCountMap[string(c)] = 1
			} else {
				characterCountMap[string(c)] += 1
			}
		}
		mapCheck[string(c)] = true
	}
}

func getWordPower(s string) int {
	power := int(0)
	mapCheck := make(map[string]bool)
	for _, c := range s {
		if mapCheck[string(c)] != true {
			power += characterCountMap[string(c)]
		}
		mapCheck[string(c)] = true
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
