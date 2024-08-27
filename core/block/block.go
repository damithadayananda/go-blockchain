package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"strings"
	"time"
)

type Block struct {
	Index        int64
	Data         interface{}
	MerkleRoot   string
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Nonce        int32
}

// Mine once the mining is completed we are returning
// and sending signal to done channel
// both the terminal operations are important if we just return only
// next mining cycle will start once we receive a new transaction only
func (b *Block) Mine(stop <-chan bool, done chan<- bool) (interrupted bool) {
	for {
		select {
		case <-stop:
			app.Logger.Info.Log("Mining is interrupted")
			// stop mining triggering new one
			done <- true
			return true
		default:
			for {
				b.Nonce++
				sha := sha256.New()
				valueString := fmt.Sprintf("merkleRoot: %v, previousHash: %v, nonce: %v ", b.MerkleRoot, b.PreviousHash, b.Nonce)
				sha.Write([]byte(valueString))
				bytString := hex.EncodeToString(sha.Sum(nil))
				subString := bytString[:3]
				numOfZeroes := strings.Count(subString, "0")
				if numOfZeroes == int(config.AppConfig.Complexity) {
					b.Hash = bytString
					// ending mining operation
					app.Logger.Info.Log("Mining is complete")
					done <- true
					return false
				}
			}
		}

	}
}

func (b *Block) CalculateMerkleRoot() {
	sha := sha256.New()
	byteArray, _ := json.Marshal(b.Data)
	sha.Write(byteArray)
	b.MerkleRoot = hex.EncodeToString(sha.Sum(nil))
}
