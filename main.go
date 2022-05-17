package main

import (
	"github.com/jackmatt2/blockchain/blockchain"
)

func main() {
	myBlockchain := blockchain.New([]byte("Genesis"))
	myBlockchain.Add([]byte("The"))
	myBlockchain.Add([]byte("fox"))
	myBlockchain.Add([]byte("jumps"))
	myBlockchain.Add([]byte("over"))
	myBlockchain.Add([]byte("the"))
	myBlockchain.Add([]byte("lazy"))
	myBlockchain.Add([]byte("dog"))
	myBlockchain.Print()
	myBlockchain.Validate()
	//myBlockchain.Blocks()[1].Data = []byte("Hacked")
	//myBlockchain.Blocks()[1].ParentHash = myBlockchain.Blocks()[3].ParentHash
	//myBlockchain.Validate()

}
