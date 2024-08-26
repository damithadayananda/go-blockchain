package blockchain

import (
	"go-blockchain/core/block"
)

type BlockChainInterface interface {
	AddBlock(block *block.Block) error
	ValidateBlock(block *block.Block) error
	GetChain() ([]block.Block, error)
	GetLastBlock() (block.Block, error)
	distributeBlock(block *block.Block) error
}
