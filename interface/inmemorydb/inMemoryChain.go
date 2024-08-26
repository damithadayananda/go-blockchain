package inmemorydb

import (
	"go-blockchain/core/block"
)

type InMemoryChain struct {
	blocks []block.Block
}

func NewInMemoryChain() *InMemoryChain {
	return &InMemoryChain{}
}
func (chain *InMemoryChain) Save(value block.Block) error {
	if len(chain.blocks) == 0 {
		value.Index = 1
		chain.blocks = append(chain.blocks, value)
		return nil
	}
	value.Index += chain.blocks[len(chain.blocks)-1].Index // increasing the index
	chain.blocks = append(chain.blocks, value)
	return nil
}

func (chain *InMemoryChain) GetAll() ([]block.Block, error) {
	return chain.blocks, nil
}

func (chain *InMemoryChain) GetLastBlock() (block.Block, error) {
	if len(chain.blocks) == 0 {
		return block.Block{}, nil
	}
	return chain.blocks[len(chain.blocks)-1], nil
}
