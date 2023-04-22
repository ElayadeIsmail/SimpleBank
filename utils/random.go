package utils

import (
	"math/rand"
	"strings"
	"time"
)

var r *rand.Rand

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Random int return int between max and Min
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

// Generate a random string of n characters
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner Generate a Random Owner string
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney Generate A random Money value between 0-1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency Generate a random Currency between USD CAD AND EUR
func RandomCurrency() string {
	currencies := []string{"USD", "CAD", "EUR"}
	return currencies[r.Intn(len(currencies))]
}
