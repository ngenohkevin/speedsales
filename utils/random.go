package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
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

func RandomWord(n int) string {
	rand.Seed(time.Now().UnixNano())

	words := make([]string, n)

	for i := 0; i < n; i++ {
		wordLength := rand.Intn(5) + 3

		word := make([]byte, wordLength)

		for j := 0; j < wordLength; j++ {
			index := rand.Intn(len(alphabet))
			word[j] = alphabet[index]
		}
		words[i] = string(word)
	}
	return strings.Join(words, " ")
}

func RandomName() string {
	return RandomString(8)
}

func RandomAddress() string {
	return RandomString(10)
}
func RandomContact() string {
	return RandomString(8)
}
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomDescription() string {
	return RandomWord(6)
}
func RandomAnyString() string {
	return RandomString(6)
}
func RandomAnyInt() int {
	return RandomInt(7, 8)
}
