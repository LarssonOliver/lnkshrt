package util

import (
	"math/rand"
	"time"

	"github.com/larssonoliver/lnkshrt/internal/config"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func NewId() string {
	return genId(config.IdLength())
}

func genId(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, length)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}
