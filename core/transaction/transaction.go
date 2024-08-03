package transaction

import (
	"github.com/google/uuid"
	"go-blockchain/controller/request"
	"go-blockchain/domain"
)

type Transaction struct {
	Id           string
	Amount       float64
	Receiver     string
	Sender       string
	Fee          float64
	Size         int
	MiningStatus domain.MiningStates
}

func (t *Transaction) generateId() {
	t.Id = uuid.New().String()
}
func (t *Transaction) generateSize() {
	// need to find the logic to get size
}

func (t *Transaction) SetMiningStatus(state domain.MiningStates) {
	t.MiningStatus = state
}

func NewTransaction(request request.TransactionRequest) Transaction {
	transaction := Transaction{
		Amount:   request.Amount,
		Receiver: request.Receiver,
		Sender:   request.Sender,
		Fee:      request.Fee,
	}
	transaction.generateId()
	transaction.generateSize()
	return transaction
}
