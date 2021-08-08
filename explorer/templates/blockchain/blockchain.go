package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevhash,omitempty"`
	Height   int    `json:"height"`
}

type blockChain struct {
	Blocks []*Block
}

var b *blockChain
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
	newBlock := Block{data, "", getLashHash(), len(GetBlockChain().Blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockChain) AddBlock(data string) {
	b.Blocks = append(b.Blocks, createBlock(data))
}

func GetBlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *blockChain) AllBlcoks() []*Block {
	return b.Blocks
}

func (b *blockChain) GetBlock(height int) *Block {
	return b.Blocks[height-1]
}
