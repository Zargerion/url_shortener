package util

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func GenerateHash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash[:8])
}

func RandomSymbol() string {
    // Диапазон ASCII-символов
    min := 33 // !
    max := 126 // ~

    r := rand.Intn(max-min+1) + min

    symbol := string(rune(r))

    return symbol
}