package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomIntBetween(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(8), RandomString(6))
}

func RandomName() string {
	return RandomString(6)
}

func RandomPassword() string {
	return RandomString(8)
}
