package response

import (
	coreBlock "go-blockchain/core/block"
	"time"
)

type BlockResponse struct {
	BaseResponse
	Result interface{} `json:"result"`
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

func ToCoreBlocks(blocks []Block) []coreBlock.Block {
	var bs []coreBlock.Block
	for _, block := range blocks {
		b := coreBlock.Block{
			Index:        block.Index,
			Data:         block.Data,
			MerkleRoot:   block.MerkleRoot,
			Hash:         block.Hash,
			PreviousHash: block.PreviousHash,
			Timestamp:    block.Timestamp,
			Nonce:        block.Nonce,
		}
		bs = append(bs, b)
	}
	return bs
}

func ToResponseBlocks(blocks []coreBlock.Block) []Block {
	var bs []Block
	for _, block := range blocks {
		b := Block{
			Index:        block.Index,
			Data:         block.Data,
			MerkleRoot:   block.MerkleRoot,
			Hash:         block.Hash,
			PreviousHash: block.PreviousHash,
			Timestamp:    block.Timestamp,
			Nonce:        block.Nonce,
		}
		bs = append(bs, b)
	}
	return bs
}
