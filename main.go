package main

import "fmt"

func main() {
	fmt.Println("main")

	wordOne, wordTwo := readWords("hi", "everyone")
}

func readWords(sOne, sTwo string) (string, string) {
	fmt.Println(sOne)
	fmt.Println(sTwo)
	return sOne, sTwo
}
