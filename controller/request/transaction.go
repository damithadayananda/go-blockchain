package request

type TransactionRequest struct {
	Amount   float64                `json:"amount"`
	Receiver string                 `json:"receiver"`
	Sender   string                 `json:"sender"`
	Fee      float64                `json:"fee"`
	Id       string                 `json:"id"` //unique id to identify transaction
	Data     map[string]interface{} `json:"data"`
}
