package main

import "sort"

var characterCountMap map[string]int = make(map[string]int)

func findAlphabetsFrequency() {
	for _, val := range currentWords {
		addWordToFreq(val)
	}
}

func addWordToFreq(s string) {
	// if a word contains the same character twice, I am counting it only once
	var mapCheck map[string]bool = make(map[string]bool)
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
	var mapCheck map[string]bool = make(map[string]bool)
	for _, c := range s {
		if mapCheck[string(c)] != true {
			power += characterCountMap[string(c)]
		}
		mapCheck[string(c)] = true

	}
	return power
}

// sorts all the words according to the probability of its coming, [NOT NEEDED]
func sortWordsByPossibility() {
	sort.Slice(sortedWords, func(i, j int) bool {
		//  think of making the sort function much more efficient via bucket sort, or choose to find the most probable word on every iteration
		return getWordPower(sortedWords[i]) > getWordPower(sortedWords[j])
	})
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
