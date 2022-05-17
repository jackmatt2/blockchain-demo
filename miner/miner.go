package miner

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// https://blockchain.info/rawblock/00000000000000000003a086c3e81956b59170c8f445f1bd9e847bba2aace9fb
// https://www.blockchain.com/btc/block/00000000000000000003a086c3e81956b59170c8f445f1bd9e847bba2aace9fb
func Mine(data []byte, difficulty int) (string, int) {
	requiredPrefix := strings.Repeat("0", difficulty)
	nonse := 0
	for {
		var dst []byte
		dst = append(dst, data...)
		dst = append(dst, []byte(fmt.Sprintf("%d", nonse))...)
		v := sha256.Sum256(dst)
		hash := fmt.Sprintf("%x", v[:])
		fmt.Printf("Attempting: %s (%d)\n", hash, nonse)
		if strings.HasPrefix(hash, requiredPrefix) {
			return hash, nonse
		}
		nonse += 1
	}
}
