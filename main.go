package main

import (
	"fmt"

	"github.com/ParkKyeongHwan/practiceCoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")
	chain.AddBlock("Fourth Block")

	for _, block := range chain.AllBlcoks() {
		fmt.Println(block.Data)
		fmt.Println(block.Hash)
		fmt.Println(block.PrevHash)
	}
}
