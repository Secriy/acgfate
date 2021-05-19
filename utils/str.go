package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const letters = "0123456789"

// GenerateCode 随机数
func GenerateCode(n int) (randStr string) {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(letters[rand.Intn(len(letters))])
	}
	randStr = string(b)
	return
}
