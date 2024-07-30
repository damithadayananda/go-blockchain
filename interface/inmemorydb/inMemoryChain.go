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
	chain.blocks = append(chain.blocks, value)
	return nil
}

func (chain *InMemoryChain) GetAll() ([]block.Block, error) {
	return chain.blocks, nil
}

func (chain *InMemoryChain) GetLastBlock() (block.Block, error) {
	return block.Block{}, nil
}
