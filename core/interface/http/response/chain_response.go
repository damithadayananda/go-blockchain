package response

import "go-blockchain/core/block"

type ChainResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"error"`
	Result  []block.Block `json:"result"`
}
