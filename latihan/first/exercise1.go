package main

import "fmt"

func main() {
	sentence := printTheResult("Saya Ahli ")
	fmt.Println(sentence)
}

func printTheResult(sentence string) string {
	newSentence := sentence + "Pemrograman Golang dan Javascript"
	return newSentence
}
