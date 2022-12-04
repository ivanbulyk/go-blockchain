package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	Nonce        int      `json:"nonce"`
	PreviousHash string   `json:"previousHash"`
	Timestamp    int64    `json:"timestamp"`
	Transactions []string `json:"transactions"`
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Timestamp:    time.Now().UnixNano(),
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp        %d\n", b.Timestamp)
	fmt.Printf("nonce        %d\n", b.Nonce)
	fmt.Printf("previousHash        %s\n", b.PreviousHash)
	fmt.Printf("transactions        %s\n", b.Transactions)
}

type Blockchain struct {
	TransactionPool []string `json:"transactionPool"`
	Chain           []*Block `json:"chain"`
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")

	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.Chain = append(bc.Chain, b)

	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.Chain {
		fmt.Printf("%s Chain  %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockChain := NewBlockchain()

	blockChain.CreateBlock(5, "hash 1")

	blockChain.CreateBlock(2, "hash 2")

	blockChain.Print()
}
