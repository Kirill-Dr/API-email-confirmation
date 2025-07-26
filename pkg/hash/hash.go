package hash

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateHash() (string, error) {
	bytes := make([]byte, 4)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
