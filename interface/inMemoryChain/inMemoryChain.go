package inMemoryChain

import "go-blockchain/core"

type InMemoryChain struct {
	blocks []core.Block
}

func (chain *InMemoryChain) Save(value core.Block) error {
	chain.blocks = append(chain.blocks, value)
	return nil
}

func (chain *InMemoryChain) GetAll() ([]core.Block, error) {
	return chain.blocks, nil
}

func (chain *InMemoryChain) GetLastBlock() (core.Block, error) {
	return core.Block{}, nil
}
