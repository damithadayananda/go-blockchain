package core

import (
	"go-blockchain/core/transaction"
)

type DataInterface interface {
	GetSize() int
}

type Transactions []transaction.Transaction

func (transactions Transactions) GetSize() int {
	return len(transactions)
}
