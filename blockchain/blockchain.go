package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

type Blockchain interface {
	Add(data []byte)
	Print()
	Validate() (bool, error)
	Blocks() []block
}

type block struct {
	Hash       string
	ParentHash string
	Data       []byte
}

type blockchain struct {
	blocks []block
}

func New(genesis []byte) Blockchain {
	return &blockchain{
		blocks: []block{{
			Hash:       toHash(genesis, ""),
			ParentHash: "",
			Data:       genesis,
		}},
	}
}

func toHash(data []byte, parentHash string) string {
	input := append(data, []byte(parentHash)...)
	hash := sha256.Sum256(input)
	return fmt.Sprintf("%x", hash)
}

func (b *blockchain) Add(data []byte) {
	parentBlock := b.blocks[len(b.blocks)-1]
	newBlock := block{
		Hash:       toHash(data, parentBlock.Hash),
		ParentHash: parentBlock.Hash,
		Data:       data,
	}
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockchain) Print() {
	for i, v := range b.blocks {
		fmt.Println("#", i)
		fmt.Println("parentHash:", v.ParentHash)
		fmt.Println("hash:      ", v.Hash)
		fmt.Println("data:      ", string(v.Data))
		fmt.Println("----------------")
	}
}

func (b *blockchain) Validate() (bool, error) {
	for i := 1; i < len(b.blocks); i++ {
		currentBlock := b.blocks[i]
		parentBlock := b.blocks[i-1]

		// Blocks must be in sequence
		if parentBlock.Hash != currentBlock.ParentHash {
			return false, errors.New("has incorrect parent hash")
		}

		// Data must match hash
		if currentBlock.Hash != toHash(currentBlock.Data, currentBlock.ParentHash) {
			return false, errors.New("has incorrect hash")
		}
	}
	return true, nil
}

func (b *blockchain) Blocks() []block {
	return b.blocks
}
