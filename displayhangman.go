package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"

	// "io/ioutil"
	"encoding/json"
	"strings"
	"time"
)

// fonction qui prend un fichier en argument et qui retourne un tableau de mots utilise
func ListeMot() string {
	dictionaryPath := os.Args[1]
	var tabword []string
	readFile, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile) // Crée un scanner pour lire le fichier.
	fileScanner.Split(bufio.ScanLines)        // Divise le fichier en lignes.
	// Parcours chaque ligne du fichier.
	for fileScanner.Scan() {
		// Ajoute le texte de la ligne actuelle à la slice trouverlemot.
		tabword = append(tabword, fileScanner.Text())
	}
	readFile.Close()
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator
	max := len(tabword)
	indexrandomword := rand.Intn(max)
	randomword := tabword[indexrandomword]

	return randomword
}

// fonction qui prend un mot en argument et qui affiche le mot en underscore avec un nombre de lettre aléatoire afficher
func Displayword(words string) string {
	numbreveal := (len(words) / 2) - 1
	tabword := make([]string, len(words))
	for i := range tabword {
		tabword[i] = "_"
	}
	for i := 0; i < numbreveal; i++ {
		max := len(words) - 1
		indexrandomword := rand.Intn(max)
		randomletter := string(words[indexrandomword])
		for j := 0; j < len(words); j++ {
			if string(words[j]) == randomletter {
				tabword[j] = randomletter
			}
		}
	}
	return strings.Join(tabword, "")
}

// fonction qui prend un mot, une lettre et un tableau de lettres en argument et qui vérifie si la lettre est dans le mot
func Imputverif(words string, letter string, tabword string) (bool, string) {
	found := false
	runes := []rune(tabword)

	for j := 0; j < len(words); j++ {
		if string(words[j]) == letter {
			runes[j] = rune(letter[0])
			found = true
		}
	}

	tabword = string(runes)
	return found, tabword
}

func Attempt(attempts int, found bool) int {
	if found {
		return attempts
	}
	return attempts - 1
}

// fonction qui prend un tableau de lettres en argument et qui vérifie si le mot est trouvé
func Win(tabword string) bool {
	for i := 0; i < len(tabword); i++ {
		if string(tabword[i]) == "_" {
			return false
		}
	}
	return true
}

// fonction qui affiche le pendu en fonction du nombre d'essai restant en lisant hangman.txt si le nombre d'essais decremente alors il ouvre le fichier hangman.txt et print le pendu correspondant

func Displaywin(win bool) {
	if win {
		fmt.Println("You win")
	} else {
		fmt.Println("You lose")
	}
}

func RemainingAttempts(maxAttempts, usedAttempts int) int {
	return maxAttempts - usedAttempts
}

func PrintHangman(attempts int) {
	// Lire les frames du fichier hangman séparées par la séquence "========="
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Error: cannot open file")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var frames []string
	var frame string     // Frame actuelle
	for scanner.Scan() { // Lit le fichier ligne par ligne
		line := scanner.Text()   // Stocke la ligne actuelle
		if line == "=========" { // Si la ligne est "=========", on a fini une frame
			frame += line + "\n"
			frames = append(frames, frame) // On ajoute la frame au tableau de frames
			frame = ""                     // On réinitialise la frame
		} else {
			frame += line + "\n" // On ajoute la ligne à la frame
		}
	}

	// Vérifiez que attempts est dans la plage valide pour éviter l'index out of range
	if attempts >= 0 && attempts < len(frames) {
		// Afficher la frame correspondant au nombre de vies restantes
		fmt.Println(frames[len(frames)-attempts-1])
	}
}

//save function
func Save(filePath string, data interface{}) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}

//load function
func Load(filePath string, data interface{}) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(data)
	if err != nil {
		return fmt.Errorf("failed to decode data: %v", err)
	}

	return nil
}

// fonction qui compare usedletter et letter et qui retourne un bool
func Compareletter(usedletter []string, letter string) bool {
	for i := 0; i < len(usedletter); i++ {
		if usedletter[i] == letter {
			return true
		}
	}
	return false
}

// fonction qui prend un mot en argument et qui retourne un tableau de lettre
func Ascii(word []rune) {

	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()

	var ascii []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ascii = append(ascii, scanner.Text())

	}

	for i := 0; i < 9; i++ {
		for j := 0; j < len(word); j++ {
			fmt.Print(ascii[((int(word[j])-32)*9)+i])
		}
		fmt.Print("\n")
	}

}
