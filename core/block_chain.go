package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go-blockchain/config"
	"go-blockchain/core/persistant"
	"strings"
	"time"
)

type BlockChain struct {
	Chain persistant.Persister
}

func NewBlockChain(chain persistant.Persister) *BlockChain {
	ch := BlockChain{
		Chain: chain,
	}
	genesisBlock := Block{
		Data:         "Genesis Block",
		PreviousHash: nil,
		Timestamp:    time.Now(),
	}
	genesisBlock.mine()
	ch.AddBlock(&genesisBlock)
	return &ch
}

func (c *BlockChain) AddBlock(block *Block) error {
	previousBlock, _ := c.Chain.GetLastBlock()
	if previousBlock.PreviousHash != nil {
		if string(previousBlock.Hash) == string(block.PreviousHash) {
			if ValidateHashComplexity(CalculateHash(*block)) {
				c.Chain.Save(*block)
			}
		}
	} else {
		c.Chain.Save(*block)
	}

	return nil
}

func CalculateHash(block Block) string {
	block.Hash = []byte{}
	by, _ := json.Marshal(block)
	sha := sha256.New()
	sha.Write(by)
	return hex.EncodeToString(sha.Sum(nil))
}

func ValidateHashComplexity(hash string) bool {
	subString := hash[:3]
	numOfZeroes := strings.Count(subString, "0")
	if numOfZeroes == int(config.AppConfig.Complexity) {
		return true
	}
	return false
}

func (c *BlockChain) ValidateBlock(block *Block) error {
	return nil
}
