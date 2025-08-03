package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/controller/request"
	"go-blockchain/core/block"
	"go-blockchain/core/interface/http/response"
	"go-blockchain/core/node"
	"go-blockchain/core/persistant"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"go-blockchain/util"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Chain BlockChainInterface

type BlockChain struct {
	database persistant.BlockChainDBInterface
}

func NewBlockChain(database persistant.BlockChainDBInterface) {
	ch := BlockChain{
		database: database,
	}
	//only genesis node will create genesis block
	if config.AppConfig.InGenesis {
		initialTransaction := transaction.NewTransaction(transaction.Transaction{
			Amount:   50, //this is the first coinbase transaction
			Receiver: app.App.Address,
		})
		genesisBlock := block.Block{
			Data:         []transaction.Transaction{initialTransaction},
			PreviousHash: "",
			Timestamp:    time.Now(),
		}
		genesisBlock.CalculateMerkleRoot()
		done := make(chan bool)
		go genesisBlock.Mine(nil, done)
		<-done
		close(done)
		ch.AddBlock([]block.Block{genesisBlock})
	} else {
		//syncing with existing blocks
		for _, node := range config.AppConfig.KnownNodes {
			client := util.GeHttpsClient(node.Certificate)
			res, err := client.Get(node.Ip + "/chain")
			if err != nil {
				log.Printf("Error in reaching Know node: %v\n", err)
				continue
			}
			data, _ := io.ReadAll(res.Body)
			var chain response.ChainResponse
			err = json.Unmarshal(data, &chain)
			if err != nil {
				continue
			}
			if err = ch.SyncBlocks(chain.Result); err != nil {
				log.Fatal("Error in syncing with main chain")
			}
			break
		}
	}
	Chain = &ch
}

func (c *BlockChain) GetChain() ([]block.Block, error) {
	return c.database.GetAll()
}

func (c *BlockChain) AddBlock(block []block.Block) error {
	for _, blk := range block {
		previousBlock, _ := c.database.GetLastBlock()
		if previousBlock.Index != 0 { // just to avoid genesis block going through the validations
			//this golden check helps to stop chain of block add calls
			//between entire cluster
			if previousBlock.Hash == blk.PreviousHash {
				// TODO need implement chain of validators here for ex:- chain length
				if c.validateHashComplexity(c.calculateHash(blk)) {
					if blk.Index > previousBlock.Index+1 {
						//this is conflict
						//need to replace last block in the existing block
						c.database.UpdateLastBlock(blk)
					}
					blk.Index = previousBlock.Index + 1
					err := c.database.Save(blk)
					if err != nil {
						return err
					}
					c.DistributeBlock(&blk)
				}
			}
		} else {
			// very first block (genesis is going through this)
			blk.Index = 1
			c.database.Save(blk)
		}
	}
	return nil
}

func (c *BlockChain) GetLastBlock() (block.Block, error) {
	return c.database.GetLastBlock()
}

func (c *BlockChain) calculateHash(block block.Block) string {
	valueString := fmt.Sprintf("merkleRoot: %v, previousHash: %v, nonce: %v ", block.MerkleRoot, block.PreviousHash, block.Nonce)
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

func (c *BlockChain) DistributeBlock(block *block.Block) error {
	nodesToBeInformed, err := c.getNodesToBeInformed()
	if err != nil {
		return err
	}

	for _, node := range nodesToBeInformed {
		reqBody := request.BlockRequest{
			Block: request.Block{
				Index:        block.Index,
				Hash:         block.Hash,
				Data:         block.Data,
				MerkleRoot:   block.MerkleRoot,
				PreviousHash: block.PreviousHash,
				Timestamp:    block.Timestamp,
				Nonce:        block.Nonce,
			},
			Metadata: map[string]interface{}{
				"caller_address": config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port),
			},
		}
		body, _ := json.Marshal(reqBody)
		req, _ := http.NewRequest(http.MethodPost, node.Ip+"/block/add", bytes.NewBuffer(body))
		client := util.GeHttpsClient(node.Certificate)
		client.Timeout = time.Second * time.Duration(config.AppConfig.BlockDistributionTimeOut)
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Add Block Request to: %s, Node: %s, Response: %v, Error: %v", string(body), node, string(data), err))
	}
	return nil
}

func (c *BlockChain) GetBlocks(noOfBlocks int) ([]block.Block, error) {
	return c.database.GetBlocks(noOfBlocks)
}

func (c *BlockChain) getNodesToBeInformed() ([]domain.Node, error) {
	knownNodes, err := node.NodeRef.GetNodes()
	if err != nil {
		return nil, err
	}
	var list []domain.Node
	for _, knownNode := range knownNodes {
		if knownNode.Ip == (config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port)) {
			continue
		}
		list = append(list, knownNode)
	}
	return list, nil
}

func (c *BlockChain) SyncBlocks(blocks []block.Block) error {
	c.database.Sync(blocks)
	return nil
}
