package helpers

import (
	"math/rand"
	"time"

	"larssonoliver.com/lnkshrt/internal/config"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func NewId() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, config.IdLength())
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
