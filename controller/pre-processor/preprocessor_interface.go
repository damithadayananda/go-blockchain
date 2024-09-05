package request

import "go-blockchain/core/block"

type IValidator interface {
	ValidateBlock(block *block.Block) error
}
