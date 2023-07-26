package main

import "fmt"

func main() {
	fmt.Println("main")

	wordOne, wordTwo := readWords("hi", "everyone")
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
