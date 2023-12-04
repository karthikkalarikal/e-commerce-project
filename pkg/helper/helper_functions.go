package helper

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomString(length int) string {
	sku := make([]byte, length)

	rand.Read(sku)

	return hex.EncodeToString(sku)
}
