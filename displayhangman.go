package hangman

import (
	"fmt"
	"math/rand"
)

func Displaywords(words string) {
	numbwords := len(words)/2 - 1

	tabword := make([]string, len(words))
	for i := range tabword {
		tabword[i] = "_"
	}

	for i := 0; i < numbwords; i++ {
		max := len(words) - 1
		indexrandomword := rand.Intn(max)
		randomletter := string(words[indexrandomword]) 

		for j := 0; j < len(words); j++ {
			if string(words[j]) == randomletter {
				tabword[j] = randomletter
			}
		}
	}
	fmt.Println(tabword)
}

func Imputuser() { 
    fmt.Println("Enter your letter: ") 
   
    var letter string 
   
    fmt.Scanln(&letter) 
}