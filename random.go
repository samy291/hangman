package main

import (
    "fmt"
    "math/rand"
)
        
func main() {
    min := 0
    max := 30
    fmt.Println(rand.Intn(max - min) + min)
} 