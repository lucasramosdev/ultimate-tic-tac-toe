package hash

import (
	"hash/fnv"
	"math/rand"
	"strings"
	"time"
)

const (
	DefaultAlphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	Min4CharsVal    int64  = 238328
	Max7CharsVal    int64  = 3521614606208
)

func CreateGameToken() *string {
	secret := "minha-chave-secreta-do-jogo"
	randomNum := generateRandomNumberInRange()
	alphabet := getShuffledAlphabet(secret)
	token := encodeBase62(randomNum, alphabet)
	return &token
}

func generateRandomNumberInRange() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	offset := Max7CharsVal - Min4CharsVal
	return r.Int63n(offset) + Min4CharsVal
}

func getShuffledAlphabet(secret string) string {
	h := fnv.New64a()
	h.Write([]byte(secret))
	seed := int64(h.Sum64())

	r := rand.New(rand.NewSource(seed))
	chars := []rune(DefaultAlphabet)

	r.Shuffle(len(chars), func(i, j int) {
		chars[i], chars[j] = chars[j], chars[i]
	})

	return string(chars)
}

func encodeBase62(num int64, alphabet string) string {
	if num == 0 {
		return string(alphabet[0])
	}

	chars := []rune(alphabet)
	base := int64(len(chars))
	var encodedBuilder strings.Builder

	for num > 0 {
		remainder := num % base
		encodedBuilder.WriteRune(chars[remainder])
		num = num / base
	}

	return reverseString(encodedBuilder.String())
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
