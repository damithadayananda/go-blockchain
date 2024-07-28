package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go-blockchain/config"
	"strings"
)

type BlockChain struct {
	Blocks []*Block
}

func (c *BlockChain) AddBlock(block *Block) error {
	if string(c.Blocks[len(c.Blocks)-1].Hash) == string(block.PreviousHash) {
		if ValidateHashComplexity(CalculateHash(*block)) {
			c.Blocks = append(c.Blocks, block)
		}
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
