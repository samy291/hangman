package random

import (
    "fmt"
    "math/rand"
)

func random() {
    min := 0
    max := 30
    fmt.Println(rand.Intn(max - min) + min)
} 