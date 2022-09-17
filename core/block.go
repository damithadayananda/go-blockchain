package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

type IBlock interface {
	mine() []byte
}

func (b *Block) mine() {
	for {
		b.Nonce++
		sha := sha256.New()
		by, _ := json.Marshal(b)
		sha.Write(by)
		bytString := hex.EncodeToString(sha.Sum(nil))
		subString := bytString[:3]
		numOfZeroes := strings.Count(subString, "0")
		if numOfZeroes == int(config.AppConfig.Complexity) {
			break
		}
	}
}
