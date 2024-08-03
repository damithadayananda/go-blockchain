package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"go-blockchain/app"
	"go-blockchain/config"
	"strings"
	"time"
)

type Block struct {
	Data         any
	Hash         []byte
	PreviousHash []byte
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
				by, _ := json.Marshal(b)
				sha.Write(by)
				bytString := hex.EncodeToString(sha.Sum(nil))
				subString := bytString[:3]
				numOfZeroes := strings.Count(subString, "0")
				if numOfZeroes == int(config.AppConfig.Complexity) {
					b.Hash = []byte(bytString)
					// ending mining operation
					app.Logger.Info.Log("Mining is complete")
					done <- true
					return false
				}
			}
		}

	}
}
