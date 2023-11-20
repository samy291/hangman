package main

import (
	"fmt"
	"hangman"
	"os"
	"strings"

	"github.com/fatih/color"
)

type HangManData struct {
	Word              string     // Word composed of '_', ex: H_ll_
	ToFind            string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts          int        // Number of attempts left
	HangmanPositions  [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	RemainingAttempts int
	Usedletter        []string
	dictionaryPath    string
	letter            string
	randletter        string
}

func main() {
	hangdat := HangManData{}

	mode := os.Args[1]

	if mode == "--classic" {

		// Use the mode and dictionary path provided as arguments
		hangdat.dictionaryPath = os.Args[2]

		// Check if --startWith flag is provided
		startWith := ""
		if len(os.Args) > 3 && strings.HasPrefix(os.Args[3], "--startWith=") {
			startWith = strings.TrimPrefix(os.Args[3], "--startWith=")
		}

		if startWith != "" {
			// Load the game from the file
			err := hangman.Load(startWith, &hangdat)
			if err != nil {
				fmt.Println("Error loading game:", err)
				return
			}
		} else {
			color.Magenta("\n			!Welcome to Hangman!")
			fmt.Println("\nYou have 10 attempts to find the word")
			hangdat.ToFind = hangman.Randomword(hangman.ListeMot(hangdat.dictionaryPath))
			hangdat.Word = hangman.Displayword(hangdat.ToFind)
		}

		// Utilisez cette variable pour suivre le nombre d'essais utilisés.
		usedAttempts := 0

		hangdat.RemainingAttempts = hangman.RemainingAttempts(10, usedAttempts)

		// Déterminez le nombre d'essais restants à partir des tentatives utilisées.
		hangdat.Attempts = hangdat.RemainingAttempts
		found := false
		// fonction save qui permet de sauvegarder la partie

	main:
		for hangdat.Attempts > 0 {
			fmt.Println(hangdat.Word)
			fmt.Println("Attempts left:", hangdat.Attempts)
			fmt.Println("Please enter a letter:")
			var letter string
			fmt.Scan(&letter)
			hangdat.letter = letter

			if letter == "STOP" {
				err := hangman.Save("save.txt", hangdat)
				if err != nil {
					fmt.Println("Error saving game:", err)
					return
				}
				fmt.Println("Game saved. You can restart with --startWith save.txt")
				break main
			}

			// Check if the letter has already been used
			if hangman.Compareletter(hangdat.Usedletter, hangdat.letter) {
				color.Red("You already used this letter")
				continue
			}

			// Vérifier si la lettre est égale au mot à trouver
			if len(letter) > 1 {
				if letter == hangdat.ToFind {
					color.Green("\nYou win")
					fmt.Println("The word was :", hangdat.ToFind)
					break
				} else {
					usedAttempts += 1
				}
			}

			// Vérifier si la lettre est dans le mot
			// Vérifier si la lettre est dans le mot
found, hangdat.Word = hangman.Imputverif(hangdat.ToFind, hangdat.letter, hangdat.Word)

			// Incrémenter le nombre d'essais utilisés uniquement si la lettre n'est pas trouvée et le rajoute à la liste des lettres utilisées
			if !found {
				usedAttempts++
				// Ajouter la lettre à la liste des lettres utilisées
				hangdat.Usedletter = append(hangdat.Usedletter, hangdat.letter)
			}

			// Mettre à jour le nombre d'essais restants en fonction du nombre d'essais utilisés
			hangdat.Attempts = hangdat.RemainingAttempts - usedAttempts

			// Afficher le pendu
			hangman.PrintHangman(hangdat.Attempts)
			hangman.Save("save.txt", hangdat)

			// Vérifier si le joueur a gagné ou perdu
			End := hangman.Win(hangdat.Word)
			if End {
				color.Green("\nYou win")
				fmt.Println("The word was :", hangdat.ToFind)
				break
			} else if hangdat.Attempts <= 0 {
				color.Red("\nYou loose")
				fmt.Println("The word was :", hangdat.ToFind)
				break
			}
		}
	}else if mode == "--ascii"{
		
		// Use the mode and dictionary path provided as arguments
		hangdat.dictionaryPath = os.Args[2]

		// Check if --startWith flag is provided
		startWith := ""
		if len(os.Args) > 3 && strings.HasPrefix(os.Args[3], "--startWith=") {
			startWith = strings.TrimPrefix(os.Args[3], "--startWith=")
		}

		if startWith != "" {
			// Load the game from the file
			err := hangman.Load(startWith, &hangdat)
			if err != nil {
				fmt.Println("Error loading game:", err)
				return
			}
		} else {
			color.Magenta("\n			!Welcome to Hangman!")
			fmt.Println("\nYou have 10 attempts to find the word")
			hangdat.ToFind = hangman.Randomword(hangman.ListeMot(hangdat.dictionaryPath))
			hangdat.Word = hangman.Displayword(hangdat.ToFind)
		}

		// Utilisez cette variable pour suivre le nombre d'essais utilisés.
		usedAttempts := 0

		hangdat.RemainingAttempts = hangman.RemainingAttempts(10, usedAttempts)

		// Déterminez le nombre d'essais restants à partir des tentatives utilisées.
		hangdat.Attempts = hangdat.RemainingAttempts
		found := false
		// fonction save qui permet de sauvegarder la partie

	asci:
		for hangdat.Attempts > 0 {
			hangman.Ascii([]rune(hangdat.Word))
			fmt.Println(hangdat.Word)
			fmt.Println("Attempts left:", hangdat.Attempts)
			fmt.Println("Please enter a letter:")
			var letter string
			fmt.Scan(&letter)
			hangdat.letter = letter

			if letter == "STOP" {
				err := hangman.Save("save.txt", hangdat)
				if err != nil {
					fmt.Println("Error saving game:", err)
					return
				}
				fmt.Println("Game saved. You can restart with --startWith save.txt")
				break asci
			}

			// Check if the letter has already been used
			if hangman.Compareletter(hangdat.Usedletter, hangdat.letter) {
				color.Red("You already used this letter")
				continue
			}

			// Vérifier si la lettre est égale au mot à trouver
			if len(letter) > 1 {
				if letter == hangdat.ToFind {
					color.Green("\nYou win")
					fmt.Println("The word was :", hangdat.ToFind)
					break
				} else {
					usedAttempts += 1
				}
			}

			// Vérifier si la lettre est dans le mot
found, hangdat.Word = hangman.Imputverif(hangdat.ToFind, hangdat.letter, hangdat.Word)

			// Incrémenter le nombre d'essais utilisés uniquement si la lettre n'est pas trouvée et le rajoute à la liste des lettres utilisées
			if !found {
				usedAttempts++
				// Ajouter la lettre à la liste des lettres utilisées
				hangdat.Usedletter = append(hangdat.Usedletter, hangdat.letter)
			}

			// Mettre à jour le nombre d'essais restants en fonction du nombre d'essais utilisés
			hangdat.Attempts = hangdat.RemainingAttempts - usedAttempts

			// Afficher le pendu
			hangman.PrintHangman(hangdat.Attempts)
			hangman.Save("save.txt", hangdat)

			// Vérifier si le joueur a gagné ou perdu
			End := hangman.Win(hangdat.Word)
			if End {
				color.Green("\nYou win")
				fmt.Println("The word was :", hangdat.ToFind)
				break
			} else if hangdat.Attempts <= 0 {
				color.Red("\nYou loose")
				fmt.Println("The word was :", hangdat.ToFind)
				break
			}
		}
	}
}
