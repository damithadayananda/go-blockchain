package util

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"go-blockchain/core"
)

func calculateHash(block core.Block) []byte {
	sha := sha256.New()
	byt, e := json.Marshal(block)
	if e != nil {
		fmt.Println(e)
	}
	sha.Write(byt)
	return sha.Sum(nil)
}
