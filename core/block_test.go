package core

import (
	"fmt"
	"go-blockchain/config"
	"testing"
	"time"
)

func TestComplexity(t *testing.T) {
	config.AppConfig.Complexity = 3
	b := Block{
		Data: "Hello World",
	}
	b.PreviousHash = []byte("000ae3b121e4b73ea0aaef435a8c439740a6cd18406605f0cc132d68d8ae947b")
	b.Timestamp = time.Now()
	b.mine()
	fmt.Println(string(b.Hash))
}
