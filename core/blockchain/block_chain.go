package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/core/block"
	"go-blockchain/core/node"
	"go-blockchain/core/persistant"
	"io"
	"net/http"
	"strings"
	"time"
)

var Chain *BlockChain

type sstring string

func (s sstring) ToString() string {
	return string(s)
}

type BlockChain struct {
	database persistant.BlockChainDBInterface
}

func NewBlockChain(database persistant.BlockChainDBInterface) {
	ch := BlockChain{
		database: database,
	}
	genesisBlock := block.Block{
		Data:         sstring("Genesis Block"),
		PreviousHash: "",
		Timestamp:    time.Now(),
	}
	done := make(chan bool)
	go genesisBlock.Mine(nil, done)
	<-done
	close(done)
	ch.AddBlock(&genesisBlock)
	Chain = &ch
}

func (c *BlockChain) GetChain() ([]block.Block, error) {
	return c.database.GetAll()
}

func (c *BlockChain) AddBlock(block *block.Block) error {
	previousBlock, _ := c.database.GetLastBlock()
	if previousBlock.Index != 0 { // just to avoid genesis block going through the validations
		if previousBlock.Hash == block.PreviousHash {
			// TODO need implement chain of validators here for ex:- chain length
			if c.validateHashComplexity(c.calculateHash(*block)) {
				block.Index += previousBlock.Index
				err := c.database.Save(*block)
				if err != nil {
					return err
				}
				c.distributeBlock(block)
			}
		}
	} else {
		// very first block (genesis is going through this)
		c.database.Save(*block)
	}

	return nil
}

func (c *BlockChain) GetLastBlock() (block.Block, error) {
	return c.database.GetLastBlock()
}

func (c *BlockChain) calculateHash(block block.Block) string {
	valueString := fmt.Sprintf("data: %v, previousHash: %v, nonce: %v ", block.Data.ToString(), block.PreviousHash, block.Nonce)
	sha := sha256.New()
	sha.Write([]byte(valueString))
	return hex.EncodeToString(sha.Sum(nil))
}

func (c *BlockChain) validateHashComplexity(hash string) bool {
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

func (c *BlockChain) distributeBlock(block *block.Block) error {
	knownNodes, err := node.NodeRef.GetNodes()
	if err != nil {
		return err
	}

	for _, node := range knownNodes {
		body, _ := json.Marshal(block)
		req, _ := http.NewRequest(http.MethodPost, node+"/block/add", bytes.NewBuffer(body))
		client := &http.Client{
			Timeout: time.Second * time.Duration(config.AppConfig.BlockDistributionTimeOut),
		}
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Request to: %s, Response: %v, Error: %v", node, string(data), err))
	}
	return nil
}
