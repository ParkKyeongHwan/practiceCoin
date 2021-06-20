package main

import (
<<<<<<< HEAD
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []block
}

func (b *blockChain) getLashHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockChain) addBlock(data string) {
	newBlock := block{data, "", b.getLashHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newBlock)
}

func (b *blockChain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("Prev Hash: %s\n", block.prevHash)
	}
}

func main() {
	chain := blockChain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
=======
	"fmt"

	"github.com/ParkKyeongHwan/practiceCoin/person"
)

func main() {
	nico := person.Person{}
	nico.SetDetails("nico", 12)
	fmt.Println("Main 'nico'", nico)
>>>>>>> 24cf20c5ee50a08d64f7659869962a884bbfcbda

}
