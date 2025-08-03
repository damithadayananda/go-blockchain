package pre_processor

import (
	"github.com/stretchr/testify/assert"
	coreBlock "go-blockchain/core/block"
	"go-blockchain/core/blockchain"
	"testing"
	"time"
)

func TestIndexValidator_ProcessBlock(t *testing.T) {
	type args struct {
		block *coreBlock.Block
	}
	tests := []struct {
		name         string
		args         args
		getLastBlock func() (coreBlock.Block, error)
		validate     func(error, []coreBlock.Block)
	}{
		{
			name: "new block is one index ahead",
			args: args{
				block: &coreBlock.Block{
					Index:     2,
					Data:      "aaa",
					Timestamp: time.Now(),
					Nonce:     5,
				}},
			validate: func(err error, blocks []coreBlock.Block) {
				assert.Nil(t, err)
			},
			getLastBlock: func() (coreBlock.Block, error) {
				return coreBlock.Block{
					Index:     1,
					Data:      "aa",
					Timestamp: time.Now(),
					Nonce:     5,
				}, nil
			},
		},
		{
			name: "new block index is equal to last block index and new one has more data",
			args: args{block: &coreBlock.Block{
				Index:     1,
				Data:      "aaa",
				Timestamp: time.Now(),
				Nonce:     5,
			}},
			getLastBlock: func() (coreBlock.Block, error) {
				return coreBlock.Block{
					Index: 1,
					Data:  "aa",
				}, nil
			},
			validate: func(err error, blocks []coreBlock.Block) {
				assert.Nil(t, err)
			},
		},
		{
			name: "new block index is one index higher than the last block index in local chain",
			args: args{block: &coreBlock.Block{
				Index:     2,
				Data:      "aaa",
				Timestamp: time.Now(),
				Nonce:     5,
			}},
			getLastBlock: func() (coreBlock.Block, error) {
				return coreBlock.Block{
					Index: 1,
					Data:  "aa",
				}, nil
			},
			validate: func(err error, blocks []coreBlock.Block) {
				assert.Nil(t, err)
				assert.Equal(t, "aaa", blocks[0].Data)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blockchain.Chain = mockChain{
				GetLastBlockFunc: tt.getLastBlock,
			}
			iv := &IndexValidator{}
			got, got1 := iv.ProcessBlock(tt.args.block, "")
			tt.validate(got, got1)
		})
	}
}

type mockChain struct {
	GetLastBlockFunc func() (coreBlock.Block, error)
}

func (mockChain) AddBlock(block []coreBlock.Block) error {
	return nil
}

func (mockChain) ValidateBlock(block *coreBlock.Block) error {
	return nil
}

func (mockChain) GetChain() ([]coreBlock.Block, error) {
	return nil, nil
}

func (m mockChain) GetLastBlock() (coreBlock.Block, error) {
	return m.GetLastBlockFunc()
}

func (mockChain) DistributeBlock(block *coreBlock.Block) error {
	return nil
}

func (mockChain) GetBlocks(noOfBlocks int) ([]coreBlock.Block, error) {
	return nil, nil
}

func (mockChain) SyncBlocks(blocks []coreBlock.Block) error {
	return nil
}
