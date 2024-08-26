package request

import (
	coreBlock "go-blockchain/core/block"
	"time"
)

type Block struct {
	Index        int64
	Data         coreBlock.DataInterface
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Nonce        int32
}

func (block Block) ToCoreBlock() coreBlock.Block {
	// instead of this manual copying we can use library like
	// github.com/jinzhu/copier
	return coreBlock.Block{
		Index:        block.Index,
		Data:         block.Data,
		Hash:         block.Hash,
		PreviousHash: block.PreviousHash,
		Timestamp:    block.Timestamp,
		Nonce:        block.Nonce,
	}
}
