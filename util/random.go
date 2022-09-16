package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
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

func RandomEmail() string {
	return fmt.Sprintf("%s@mail.com", RandomString(6))
}

func RandomUnitNo() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var sb strings.Builder
	sb.WriteString(RandomString(1))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(rand.Intn(99)))
	sb.WriteString("-")
	sb.WriteString(strconv.Itoa(rand.Intn(9999)))
	return sb.String()
}

func RandomPhone() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var sb strings.Builder
	sb.WriteString("628")
	for i := 0; i < 9; i++ {
		sb.WriteString(strconv.Itoa(rand.Intn(9)))
	}
	return sb.String()
}
