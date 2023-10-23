package main

import (
	"hangman"
)

func main() {
	tabWord := hangman.ExtractWordsFromFile("Words/words.txt")
	randomWord := hangman.Random(tabWord)
	hangman.Displaywords(randomWord)
}
