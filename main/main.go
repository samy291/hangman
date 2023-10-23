package main

import (
	"hangman"
)

func main() {
	tabWord := hangman.ExtractWordsFromFile("Words/words.txt")
	randomWord := hangman.Random(tabWord)
	tabletter := hangman.Displaywords(randomWord)
	letterimput := hangman.Imputuser()
	hangman.Imputverif(randomWord, letterimput, tabletter)
}
