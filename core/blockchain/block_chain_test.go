package blockchain

import (
	"github.com/stretchr/testify/assert"
	"go-blockchain/config"
	"go-blockchain/core/block"
	"go-blockchain/core/transaction"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	config.AppConfig.Complexity = 3
	bc := BlockChain{}
	b := block.Block{
		Index: 0,
		Data: transaction.List([]transaction.Transaction{
			transaction.Transaction{
				Id:           "504e914d-0b23-4091-a8d4-ba047cc67cc9",
				Amount:       300,
				Receiver:     "287335e156a4d7b5e1252af50c6c5592286793f8a78a9c8d6eb69e9f1974e41c",
				Sender:       "c7f243735dbb5eb0b9899567f0784dfc3684b1bdd8d7ce1500b3c56b337d4b8d",
				Fee:          5,
				Size:         0,
				MiningStatus: 0,
			},
			transaction.Transaction{
				Id:           "38dc2294-f940-47e9-bec5-2c316d1ffa2a",
				Amount:       300,
				Receiver:     "287335e156a4d7b5e1252af50c6c5592286793f8a78a9c8d6eb69e9f1974e41c",
				Sender:       "c7f243735dbb5eb0b9899567f0784dfc3684b1bdd8d7ce1500b3c56b337d4b8d",
				Fee:          5,
				Size:         0,
				MiningStatus: 0,
			},
		}),
		PreviousHash: "0001bf18bbd4a96f764a2b4d91a9e5923d9510ace2a595150c374f0133efbdab",
		Nonce:        6166,
	}
	valid := bc.validateHashComplexity(bc.calculateHash(b))
	assert.True(t, valid)
}
