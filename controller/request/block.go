package request

import (
	"errors"
	coreBlock "go-blockchain/core/block"
	"time"
)

type BlockRequest struct {
	Block    Block                  `json:"block"`
	Metadata map[string]interface{} `json:"metadata"`
}
type Block struct {
	Index        int64       `json:"index"`
	Data         interface{} `json:"data"`
	MerkleRoot   string      `json:"merkle_root"`
	Hash         string      `json:"hash"`
	PreviousHash string      `json:"previous_hash"`
	Timestamp    time.Time   `json:"timestamp"`
	Nonce        int32       `json:"nonce"`
}

func (blockRq BlockRequest) ToCoreBlock() (coreBlock.Block, error) {
	//TODO here to support multiple data types, future we will implement Type registry
	//data, ok := block.Data.(transaction.List)
	//if ok {
	//	app.Logger.Error.Log("UnMarshaling error: %v", block.Data)
	//}
	// instead of this manual copying we can use library like
	// github.com/jinzhu/copier
	return coreBlock.Block{
		Index:        blockRq.Block.Index,
		Data:         blockRq.Block.Data,
		Hash:         blockRq.Block.Hash,
		PreviousHash: blockRq.Block.PreviousHash,
		Timestamp:    blockRq.Block.Timestamp,
		Nonce:        blockRq.Block.Nonce,
		MerkleRoot:   blockRq.Block.MerkleRoot,
	}, nil
}

func (blockRq BlockRequest) Validate() error {
	if blockRq.Metadata == nil {
		return errors.New("block metadata is required")
	}
	return nil
}
