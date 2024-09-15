package transaction

import (
	"github.com/google/uuid"
	"go-blockchain/domain"
)

type Transaction struct {
	Id           string              `json:"id"`
	Amount       float64             `json:"amount"`
	Receiver     string              `json:"receiver"`
	Sender       string              `json:"sender"`
	Fee          float64             `json:"fee"`
	Size         int                 `json:"size"`
	MiningStatus domain.MiningStates `json:"miningStatus"`
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

func NewTransaction(request Transaction) Transaction {
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
