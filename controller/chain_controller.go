package controller

import (
	"go-blockchain/controller/response"
	"go-blockchain/core/blockchain"
)

type ChainController interface {
	GetChain() response.ChainResponse
}

type ChainControllerImpl struct {
}

func (cr *ChainControllerImpl) GetChain() response.ChainResponse {
	chain, _ := blockchain.Chain.GetChain()
	return response.ChainResponse{
		SuccessResponse: response.SuccessResponse{
			BaseResponse: response.BaseResponse{
				Success: true,
			},
		},
		Result: chain,
	}
}
