package miner

import (
	"fmt"
	"testing"
)

func TestMine(t *testing.T) {
	hash, nonse := Mine([]byte("Hello World"), 1)
	fmt.Printf("Found: %s (%d)", hash, nonse)
}
