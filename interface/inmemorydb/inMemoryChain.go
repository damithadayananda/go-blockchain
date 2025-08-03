package inmemorydb

import (
	"errors"
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
	if len(chain.blocks) == 0 {
		return block.Block{}, nil
	}
	return chain.blocks[len(chain.blocks)-1], nil
}

func (chain *InMemoryChain) GetBlocks(numOfBlocks int) ([]block.Block, error) {
	if len(chain.blocks) < numOfBlocks {
		return nil, errors.New("chain length is less than the number of requested blocks")
	}
	return chain.blocks[:numOfBlocks], nil
}

func (chain *InMemoryChain) UpdateLastBlock(block block.Block) {
	chain.blocks[len(chain.blocks)-1] = block
}

func (chain *InMemoryChain) Sync(blocks []block.Block) error {
	chain.blocks = append(chain.blocks, blocks...)
	return nil
}
