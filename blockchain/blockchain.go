package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type BlockChain struct {
	Blocks []*Block
}

var b *BlockChain
var once sync.Once

func getLashHash() string {
	totalBlocks := len(GetBlockChain().Blocks)
	if totalBlocks == 0 {
		return ""
	}

	return GetBlockChain().Blocks[totalBlocks-1].Hash
}

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLashHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *BlockChain) AddBlock(data string) {
	b.Blocks = append(b.Blocks, createBlock(data))
}

func GetBlockChain() *BlockChain {
	if b == nil {
		once.Do(func() {
			b = &BlockChain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *BlockChain) AllBlcoks() []*Block {
	return b.Blocks
}
