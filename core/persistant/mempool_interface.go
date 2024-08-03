package persistant

import (
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
)

type MemPoolInterface interface {
	Save(transaction transaction.Transaction) error
	Get() ([]transaction.Transaction, error)
	GetAll() []transaction.Transaction
	Delete(in string)
	MiningStatusUpdate(id string, status domain.MiningStates) error
}
