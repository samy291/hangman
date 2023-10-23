package hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func ExtractWordsFromFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Erreur lors de la lecture du fichier", err)
		os.Exit(1)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Erreur lors de la lecture du fichier", err)
		os.Exit(1)
	}
	return words
}

func Random(words []string) string {
	max := len(words) - 1
	indexrandom := (rand.Intn(max))
	randomwords := words[indexrandom]
	fmt.Println(randomwords)
	return randomwords
}
