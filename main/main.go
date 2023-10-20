package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

func main() {
    directoryPath := "Words" 

    dir, err := os.Open(directoryPath)
    if err != nil {
        log.Fatal(err)
    }
    defer dir.Close()

    files, err := dir.Readdir(-1)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if file.Mode().IsRegular() {
            filePath := filepath.Join(directoryPath, file.Name())
            words, err := extractWordFile(filePath)
            if err != nil {
                log.Printf("Erreur lors de la lecture du fichier %s : %v\n", filePath, err)
                continue
            }
            fmt.Println(strings.Join(words, "\n"))
        }
    }
}

func extractWordFile(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var words []string
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return words, nil
}
