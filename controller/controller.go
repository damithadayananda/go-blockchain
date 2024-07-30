package controller

import (
	"encoding/json"
	request2 "go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/blockchain"
	"go-blockchain/core/mempool"
	"go-blockchain/core/transaction"
	"net/http"
)

type ApiController interface {
	GetChain() response.ChainResponse
	AddTransaction(request http.Request) error
}

type ApiControllerImpl struct {
}

func (cr *ApiControllerImpl) GetChain() response.ChainResponse {
	chain, _ := blockchain.Chain.GetChain()
	return response.ChainResponse{
		SuccessResponse: response.SuccessResponse{
			Success: true,
		},
		Result: chain,
	}
}

func (cr *ApiControllerImpl) AddTransaction(r *http.Request) error {
	txnRequest := request2.TransactionRequest{}
	if err := json.NewDecoder(r.Body).Decode(&txnRequest); err != nil {
		return err
	}
	txn := transaction.NewTransaction(txnRequest)
	mempool.Mempool.Save(txn)
	return nil
}
