package hangman

import (
	"fmt"
	"math/rand"
)

func Displaywords(words string) []string {
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
	return tabword
}

func Imputuser(numattempt int) string {
	if numattempt >= 1 {
		fmt.Println("Enter your letter: ")

		var letter string

		fmt.Scanln(&letter)
		return letter
	}
	return ""
}

func Imputverif(words string, letter string, tabword []string) bool {
	found := false

	for j := 0; j < len(words); j++ {
		if string(words[j]) == letter {
			tabword[j] = letter
			found = true
		}
	}
	fmt.Println(tabword)
	return found
}

func Attempt(bool bool) int {
	attempts := 10
	if !bool {
		attempts--
	}
	fmt.Println(attempts)
	return attempts
}
