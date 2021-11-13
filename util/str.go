package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

const (
	numLetters  = "0123456789"
	charLetters = "0123456789abcdefghijklmnopqrstuvwxyz"
)

// GenerateCode 随机数
func GenerateCode(n int) (randStr string) {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(numLetters[rand.Intn(len(numLetters))])
	}
	randStr = string(b)
	return
}

// GenerateStr 随机字符串
func GenerateStr(n int) (randStr string) {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(charLetters[rand.Intn(len(charLetters))])
	}
	randStr = string(b)
	return
}
