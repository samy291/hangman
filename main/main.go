package main

import (
	"fmt"
	"hangman"

	"github.com/fatih/color"
)

type HangManData struct {
	Word              []string   // Word composed of '_', ex: H_ll_
	ToFind            string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts          int        // Number of attempts left
	HangmanPositions  [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	RemainingAttempts int
	Usedletter        []string
}

// fonction main qui utilise les structures HangManData et qui appelle les fonctions du package hangman sans utiliser hangman.play()
func main() {
	hangdat := HangManData{}
	color.Magenta("\n			!Welcome to Hangman!")
	fmt.Println("\nYou have 10 attempts to find the word")
	hangdat.ToFind = hangman.Randomword(hangman.ListeMot("Words/words.txt"))
	hangdat.Word, hangdat.Usedletter = hangman.Displayword(hangdat.ToFind)

	// Utilisez cette variable pour suivre le nombre d'essais utilisés.
	usedAttempts := 0

	hangdat.RemainingAttempts = hangman.RemainingAttempts(10, usedAttempts)

	// Déterminez le nombre d'essais restants à partir des tentatives utilisées.
	hangdat.Attempts = hangdat.RemainingAttempts
	found := false

	for hangdat.Attempts > 0 {
		fmt.Println(hangdat.Word)
		fmt.Println("Attempts left:", hangdat.Attempts)
		fmt.Println("Please enter a letter:")
		var letter string
		fmt.Scan(&letter)

		// Vérifier si la lettre est égale au mot à trouver
		// if len(letter) > 1 {
		// if letter == hangdat.ToFind {
		// 	color.Green("\nYou win")
		// 	fmt.Println("The word was :", hangdat.ToFind)
		// 	break
		// }else{
		// 	usedAttempts+=2
		// }
		// }
		//Vérifier si la lettre a déjà été utilisée
		for i := 0; i < len(hangdat.Usedletter); i++ {
			if letter == hangdat.Usedletter[i] {
				color.Red("\nYou already used this letter")
				found = true
			}
		}
		// Vérifier si la lettre est dans le mot
		found, hangdat.Usedletter = hangman.Imputverif(hangdat.ToFind, letter, hangdat.Word)

		// Incrémenter le nombre d'essais utilisés uniquement si la lettre n'est pas trouvée
		if !found {
			usedAttempts++
		}

		// Mettre à jour le nombre d'essais restants en fonction du nombre d'essais utilisés
		hangdat.Attempts = hangdat.RemainingAttempts - usedAttempts

		// Afficher le pendu
		hangman.PrintHangman(hangdat.Attempts)

		// Vérifier si le joueur a gagné ou perdu
		End := hangman.Win(hangdat.Word)
		if End {
			color.Green("\nYou win")
			fmt.Println("The word was :", hangdat.ToFind)
			break
		} else if hangdat.Attempts == 0 {
			color.Red("\nYou lose")
			fmt.Println("The word was :", hangdat.ToFind)
			break
		}
	}
}
