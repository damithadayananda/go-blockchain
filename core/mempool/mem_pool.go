package mempool

import (
	"go-blockchain/core/persistant"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
)

var Mempool *MemPool

type MemPool struct {
	database persistant.MemPoolInterface
	Receiver chan bool
	Remover  chan []string
}

func NewMemPool(database persistant.MemPoolInterface) *MemPool {
	Mempool = &MemPool{
		Receiver: make(chan bool),
		Remover:  make(chan []string),
		database: database,
	}
	return Mempool
}

func (m MemPool) Save(transaction transaction.Transaction) error {
	err := m.database.Save(transaction)
	if err != nil {
		return err
	}
	m.Receiver <- true
	return nil
}

func (m MemPool) Get() ([]transaction.Transaction, error) {
	return m.database.Get()
}

func (m MemPool) Delete(ids []string) error {
	for _, id := range ids {
		m.database.Delete(id)
	}
	m.Remover <- ids
	return nil
}

func (m MemPool) Mark(transactions []transaction.Transaction) error {
	for _, txn := range transactions {
		m.database.MiningStatusUpdate(txn.Id, domain.MINING_DONE)
	}
	return nil
}
