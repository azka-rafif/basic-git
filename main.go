package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("main")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("word one")
	scanner.Scan()
	fmt.Println("word two")
	wordOne := scanner.Text()
	scanner.Scan()
	wordTwo := scanner.Text()
	wordOne, wordTwo = readWords(wordOne, wordTwo)
	sentence := turnIntoSentence(wordOne, wordTwo)
	printALot(10, sentence)
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
