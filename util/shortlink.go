package util

import (
	"math/rand"
	"time"
)

const shortLinkLength = 4

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// GenerateShortLink generates a short link
func GenerateShortLink() string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, shortLinkLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
