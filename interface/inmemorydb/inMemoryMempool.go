package inmemorydb

import (
	"go-blockchain/config"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"sort"
	"sync"
)

type InMemoryMemPool struct {
	mempool []*transaction.Transaction
	lock    sync.RWMutex
}

func NewInMemoryMemPool() *InMemoryMemPool {
	return &InMemoryMemPool{}
}
func (im *InMemoryMemPool) Save(transaction transaction.Transaction) error {
	im.lock.Lock()
	defer im.lock.Unlock()
	im.mempool = append(im.mempool, &transaction)
	return nil
}

// TODO need to think of more generic implementation
func (im *InMemoryMemPool) Get() ([]transaction.Transaction, error) {
	im.lock.RLock()
	defer im.lock.RUnlock()
	var eligibleTransactions []transaction.Transaction
	for _, txn := range im.mempool {
		if txn.MiningStatus == domain.READY_FOR_MINING {
			eligibleTransactions = append(eligibleTransactions, *txn)
		}
	}
	sort.Slice(eligibleTransactions, func(i, j int) bool {
		return eligibleTransactions[i].Fee >= eligibleTransactions[j].Fee
	})
	if len(eligibleTransactions) >= config.AppConfig.MaxTransactionsPerBlock {
		return eligibleTransactions[:config.AppConfig.MaxTransactionsPerBlock], nil
	}
	return nil, nil
}

func (im *InMemoryMemPool) GetAll() []transaction.Transaction {
	im.lock.RLock()
	defer im.lock.RUnlock()
	var transaction []transaction.Transaction
	for _, txn := range im.mempool {
		transaction = append(transaction, *txn)
	}
	return transaction
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

func (im *InMemoryMemPool) MiningStatusUpdate(id string, status domain.MiningStates) error {
	im.lock.Lock()
	defer im.lock.Unlock()
	for _, v := range im.mempool {
		if v.Id == id {
			v.MiningStatus = status
		}
	}
	return nil
}
