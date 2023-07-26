package main

import "fmt"

func main() {
	fmt.Println("main")

	wordOne, wordTwo := readWords("hi", "everyone")
	sentence := turnIntoSentence(wordOne, wordTwo)
}

func readWords(sOne, sTwo string) (string, string) {
	fmt.Println(sOne)
	fmt.Println(sTwo)
	return sOne, sTwo
}

func turnIntoSentence(wOne, wTwo string) string {
	fmt.Printf("%v %v", wOne, wTwo)
	return fmt.Sprintf("%v %v", wOne, wTwo)
}
