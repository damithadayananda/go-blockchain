package util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go-blockchain/core/block"
)

func calculateHash(block block.Block) []byte {
	sha := sha256.New()
	byt, e := json.Marshal(block)
	if e != nil {
		fmt.Println(e)
	}
	sha.Write(byt)
	return sha.Sum(nil)
}
