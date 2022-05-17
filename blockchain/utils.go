package blockchain

import (
	"crypto/sha256"
	"fmt"
)

func toHash(data []byte) string {
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}
