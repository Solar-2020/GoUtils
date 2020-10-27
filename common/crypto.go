package common

import (
	"math/rand"
	"strings"
	"time"
)

const (
	Capitals = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowers = "abcdefghijklmnopqrstuvwxyz"
	Numbers = "0123456789"
	LightSymbols = "_-"
	HardSymbols = `,.!@#$%^&*(){}[]+="';'/`
)

func GetRandomString(length int64, alphabets ...string) string {
	return GetRandomStringWithSeed(length, time.Now().UnixNano(), alphabets...)
}

func GetRandomStringWithSeed(length, seed int64, alphabets ...string) string {
	chars := []rune(strings.Join(alphabets, ""))
	rand.Seed(seed)
	var b strings.Builder
	for i := int64(0); i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
