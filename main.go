package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("main")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("word one")
	scanner.Scan()
	wordOne := scanner.Text()

	fmt.Println("word two")
	scanner.Scan()
	wordTwo := scanner.Text()

	wordOne, wordTwo = readWords(wordOne, wordTwo)
	sentence := turnIntoSentence(wordOne, wordTwo)

	fmt.Println("Enter number of times to print: ")
	scanner.Scan()
	times, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printALot(times, sentence)
}

func readWords(sOne, sTwo string) (string, string) {
	fmt.Println(sOne)
	fmt.Println(sTwo)
	return sOne, sTwo
}

func turnIntoSentence(wOne, wTwo string) string {
	fmt.Printf("%v %v\n", wOne, wTwo)
	return fmt.Sprintf("%v %v", wOne, wTwo)
}

func printALot(times int, sentence string) {
	for i := 0; i < times; i++ {
		fmt.Println(sentence)
	}
}
