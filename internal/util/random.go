package util

import (
	"math/rand"
	"strings"
	"time"
)

const NumCharset = "0123456789"
const LAlphaCharset = "abcdefghijklmnopqrstuvwxyz"
const CAlphaCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(length int, charsets ...string) string {
	charset := strings.Join(charsets, "")

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
