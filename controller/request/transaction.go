package request

type TransactionRequest struct {
	Amount   float64 `json:"amount"`
	Receiver string  `json:"receiver"`
	Sender   string  `json:"sender"`
	Fee      float64 `json:"fee"`
}
