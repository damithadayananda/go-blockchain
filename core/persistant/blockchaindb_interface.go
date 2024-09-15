package persistant

import (
	"go-blockchain/core/block"
)

type BlockChainDBInterface interface {
	Save(value block.Block) error
	GetAll() ([]block.Block, error)
	GetLastBlock() (block.Block, error)
	GetBlocks(numOfBlocks int) ([]block.Block, error)
	UpdateLastBlock(block block.Block)
}
