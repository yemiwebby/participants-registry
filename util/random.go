package util

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/google/uuid"
)


const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomReferenceNumber() string {
	let := RandomString(3)
	num := RandomInt(1, 100)

	return strings.Join([]string{strconv.Itoa(int(num)), let}, "")
}

func RandomString(n int) string {
	var sb strings.Builder
	number := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn((number))]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}