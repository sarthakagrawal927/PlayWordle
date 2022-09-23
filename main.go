package main

func main() {
	// Code to play the game
	// var status string
	// var currWord string
	// reader := bufio.NewReader(os.Stdin)
	// currWordCount := int(0)
	// copy(currentWords, wordList)
	// for status != "WIN" {
	// 	findAlphabetsFrequency()
	// 	currWord = getMostProbableWord(currWordCount)
	// 	fmt.Println(currWord, len(currentWords), characterCountMap)
	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("An error occurred while reading input. Please try again", err)
	// 		return
	// 	}
	// 	status = strings.ToUpper(strings.TrimSuffix(input, "\n"))
	// 	moves := makePlayObject(currWord, status)
	// 	currWordCount += 1
	// 	playMoves(moves)
	// }

	// Code to simulate the game

	// fmt.Println(simulate("APPLE"), len(currentWords), len(wordList))
	// fmt.Println(simulate("ADZED"), len(currentWords), len(wordList))
	// fmt.Println(simulate("HEVEA"), len(currentWords), len(wordList))
	// fmt.Println(simulate("YRIVD"), len(currentWords), len(wordList))
	// moves := make(chan int, 1)
	// simulate("ZULUS", moves)
	// fmt.Println(<-moves)
	simulateForAll()

	// fmt.Println(currentWords)
}

var showDebug bool = true

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
