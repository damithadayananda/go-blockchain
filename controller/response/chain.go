package response

import (
	"go-blockchain/core/block"
)

type ChainResponse struct {
	SuccessResponse
	Result []block.Block `json:"result"`
}
