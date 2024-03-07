package util

import (
	"crypto/sha256"
	"math/big"
	"math/rand"
)

func GenerateHash(s string) string {
	hash := sha256.Sum256([]byte(s))

	// Преобразуем хеш в число
	hashInt := new(big.Int)
	hashInt.SetBytes(hash[:])

	// Получаем хеш в 36-ричной системе счисления (цифры + латинские буквы)
	hashString := hashInt.Text(36)
	return hashString[:7]
}

func RandomSymbol() string {
	// Диапазон ASCII-символов
	min := 33  // !
	max := 126 // ~

	r := rand.Intn(max-min+1) + min

	symbol := string(rune(r))

	return symbol
}
