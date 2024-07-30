package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go-blockchain/config"
	"go-blockchain/core/block"
	"go-blockchain/core/persistant"
	"strings"
	"time"
)

var Chain *BlockChain

type BlockChain struct {
	database persistant.BlockChainDBInterface
}

func NewBlockChain(database persistant.BlockChainDBInterface) {
	ch := BlockChain{
		database: database,
	}
	genesisBlock := block.Block{
		Data:         "Genesis Block",
		PreviousHash: nil,
		Timestamp:    time.Now(),
	}
	genesisBlock.Mine()
	ch.AddBlock(&genesisBlock)
	Chain = &ch
}

func (c *BlockChain) GetChain() ([]block.Block, error) {
	return c.database.GetAll()
}

func (c *BlockChain) AddBlock(block *block.Block) error {
	previousBlock, _ := c.database.GetLastBlock()
	if previousBlock.PreviousHash != nil {
		if string(previousBlock.Hash) == string(block.PreviousHash) {
			if ValidateHashComplexity(CalculateHash(*block)) {
				c.database.Save(*block)
			}
		}
	} else {
		c.database.Save(*block)
	}

	return nil
}

func CalculateHash(block block.Block) string {
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

func (c *BlockChain) ValidateBlock(block *block.Block) error {
	return nil
}
