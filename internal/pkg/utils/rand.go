package utils

import (
	"math/rand"
	"time"
)

var SeededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(n int) string {
	const letters = "0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[SeededRand.Intn(len(letters))]
	}
	return string(b)
}
