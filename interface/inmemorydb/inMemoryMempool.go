package inmemorydb

import (
	"go-blockchain/config"
	"go-blockchain/core/transaction"
	"sort"
	"sync"
)

type InMemoryMemPool struct {
	mempool []transaction.Transaction
	lock    sync.RWMutex
}

func NewInMemoryMemPool() *InMemoryMemPool {
	return &InMemoryMemPool{}
}
func (im *InMemoryMemPool) Save(transaction transaction.Transaction) error {
	im.lock.Lock()
	defer im.lock.Unlock()
	im.mempool = append(im.mempool, transaction)
	return nil
}

func (im *InMemoryMemPool) Get() ([]transaction.Transaction, error) {
	im.lock.RLock()
	defer im.lock.RUnlock()
	sort.Slice(im.mempool, func(i, j int) bool {
		return im.mempool[i].Fee >= im.mempool[j].Fee
	})
	if len(im.mempool) >= config.AppConfig.MaxTransactionsPerBlock {

		return im.mempool[:config.AppConfig.MaxTransactionsPerBlock], nil
	}
	return nil, nil
}

func (im *InMemoryMemPool) GetAll() []transaction.Transaction {
	im.lock.RLock()
	defer im.lock.RUnlock()
	return im.mempool
}

func (im *InMemoryMemPool) Delete(id string) {
	im.lock.Lock()
	defer im.lock.Unlock()
	for k, v := range im.mempool {
		if v.Id == id {
			im.mempool = append(im.mempool[:k], im.mempool[k+1:]...)
		}
	}
}
