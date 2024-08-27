package request

import (
	coreBlock "go-blockchain/core/block"
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

func (block Block) ToCoreBlock() (coreBlock.Block, error) {
	//TODO here to support multiple data types, future we will implement Type registry
	//data, ok := block.Data.(transaction.List)
	//if ok {
	//	app.Logger.Error.Log("UnMarshaling error: %v", block.Data)
	//}
	// instead of this manual copying we can use library like
	// github.com/jinzhu/copier
	return coreBlock.Block{
		Index:        block.Index,
		Data:         block.Data,
		Hash:         block.Hash,
		PreviousHash: block.PreviousHash,
		Timestamp:    block.Timestamp,
		Nonce:        block.Nonce,
		MerkleRoot:   block.MerkleRoot,
	}, nil
}
