package utils

import (
	"crypto/sha256"
	"fmt"
)

func GenerateHashPassword(email, password string) string {
	hashStr := fmt.Sprintf("%s:%s", email, password)

	h := sha256.New()
	h.Write([]byte(hashStr))

	return string(h.Sum(nil))
}
