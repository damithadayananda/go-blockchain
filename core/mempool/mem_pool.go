package mempool

import (
	"go-blockchain/core/persistant"
	"go-blockchain/core/transaction"
)

var Mempool *MemPool

type MemPool struct {
	database persistant.MemPoolInterface
}

func NewMemPool(database persistant.MemPoolInterface) {
	Mempool = &MemPool{
		database: database,
	}
}

func (m MemPool) Save(transaction transaction.Transaction) error {
	return m.database.Save(transaction)
}

func (m MemPool) Get() ([]transaction.Transaction, error) {
	return m.database.Get()
}
