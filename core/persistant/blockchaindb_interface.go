package persistant

import (
	"go-blockchain/core/block"
)

type BlockChainDBInterface interface {
	Save(value block.Block) error
	GetAll() ([]block.Block, error)
	GetLastBlock() (block.Block, error)
}
