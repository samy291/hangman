package main

import (

	"flag"
	"fmt"
	"hangman"
	"os"
	"strings"

	"github.com/fatih/color"
)

type HangManData struct {
	Word              []string   // Word composed of '_', ex: H_ll_
	ToFind            string     // Final word chosen by the program at the beginning. It is the word to find
	Attempts          int        // Number of attempts left
	HangmanPositions  [10]string // It can be the array where the positions parsed in "hangman.txt" are stored
	RemainingAttempts int
	Usedletter        []string
	dictionaryPath    string
} 

func main() {
	hangdat := HangManData{}
    if len(os.Args) < 3 || (os.Args[1] != "--classic" && os.Args[1] != "asciiart") {
        fmt.Println("Usage: ./main/hangman.go --classic <dictionary> or ./main/hangman.go asciiart <dictionary>")
        return
    }

    // Use the mode and dictionary path provided as arguments
    mode := os.Args[1]
    hangdat.dictionaryPath = os.Args[2]

    // Check if --startWith flag is provided
    startWith := ""
    if len(os.Args) > 3 && strings.HasPrefix(os.Args[3], "--startWith=") {
        startWith = strings.TrimPrefix(os.Args[3], "--startWith=")
    }

    if startWith != "" {
        // Load saved data
        if err := hangman.Load(startWith, &hangdat); err != nil {
            fmt.Println("Error loading saved data:", err)
            return
        }
    } else {
        color.Magenta("\n			!Welcome to Hangman!")
        fmt.Println("\nYou have 10 attempts to find the word")
        hangdat.ToFind = hangman.Randomword(hangman.ListeMot(hangdat.dictionaryPath))
        hangdat.Word, hangdat.Usedletter = hangman.Displayword(hangdat.ToFind)
    }


	// Utilisez cette variable pour suivre le nombre d'essais utilisés.
	usedAttempts := 0

	hangdat.RemainingAttempts = hangman.RemainingAttempts(10, usedAttempts)

	// Déterminez le nombre d'essais restants à partir des tentatives utilisées.
	hangdat.Attempts = hangdat.RemainingAttempts
	found := false
	// fonction save qui permet de sauvegarder la partie

mainLoop:
	for hangdat.Attempts > 0 {
		fmt.Println(hangdat.Word)
		fmt.Println("Attempts left:", hangdat.Attempts)
		fmt.Println("Used letters:", hangdat.Usedletter)
		fmt.Println("Please enter a letter:")
		var letter string
		fmt.Scan(&letter)

		switch strings.ToUpper(letter) {
		case "STOP":
			hangman.Save("save.txt", hangdat.Word, hangdat.Attempts, hangdat.Usedletter)
			fmt.Println("Game saved. You can restart with --startWith save.txt")
			break mainLoop
		}

		// Check if the letter has already been used
		Letterused := hangman.IsLetterUsed(letter, hangdat.Usedletter)
		
		if Letterused == true{
			fmt.Println("This letter has already been used")
			continue
		}

		// Mark the letter as used
		hangdat.Usedletter = append(hangdat.Usedletter, letter)

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
		found, hangdat.Usedletter = hangman.Imputverif(hangdat.ToFind, letter, hangdat.Word)

		// Incrémenter le nombre d'essais utilisés uniquement si la lettre n'est pas trouvée
		if !found {
			usedAttempts++
		}

		// Mettre à jour le nombre d'essais restants en fonction du nombre d'essais utilisés
		hangdat.Attempts = hangdat.RemainingAttempts - usedAttempts

		// Afficher le pendu
		hangman.PrintHangman(hangdat.Attempts)
		hangman.Save("save.txt", hangdat.Word, hangdat.Attempts, hangdat.Usedletter)

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
}