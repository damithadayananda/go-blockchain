package response

import "go-blockchain/core/transaction"

type TransactionResponse struct {
	BaseResponse
	Result []Transaction
}

type Transaction struct {
	Id           string                 `json:"id"`
	Amount       float64                `json:"amount"`
	Receiver     string                 `json:"receiver"`
	Sender       string                 `json:"sender"`
	Fee          float64                `json:"fee"`
	Size         int                    `json:"size"`
	MiningStatus string                 `json:"miningStatus"`
	Data         map[string]interface{} `json:"data"`
}

func (t *Transaction) FromDomain(transaction transaction.Transaction) {
	t.Id = transaction.Id
	t.Amount = transaction.Amount
	t.Receiver = transaction.Receiver
	t.Sender = transaction.Sender
	t.Fee = transaction.Fee
	t.Size = transaction.Size
	t.MiningStatus = transaction.MiningStatus.String()
	t.Data = transaction.Data
}
