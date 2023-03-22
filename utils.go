package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readWordFromCli(reader bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading input. Please try again", err)
		os.Exit(1)
	}
	return input
}

func getFormattedString(input string) string {
	return strings.ToUpper(strings.TrimSuffix(input, "\n"))
}

func splitInput(input string) (string, string) {
	temp := strings.Split(input, " ")
	temp[0], temp[1] = getFormattedString(temp[0]), getFormattedString(temp[1])
	if len(temp) != 2 || (len(temp[0]) != 5 && temp[0] != "TEST") || len(temp[1]) != 5 {
		fmt.Println("Please enter a valid input")
		os.Exit(1)
	}
	return temp[0], temp[1]
}
