package persistant

import "go-blockchain/core/transaction"

type MemPoolInterface interface {
	Save(transaction transaction.Transaction) error
	Get() ([]transaction.Transaction, error)
	GetAll() []transaction.Transaction
	Delete(in string)
}
