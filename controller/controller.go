package controller

import (
	"encoding/json"
	"go-blockchain/app"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/blockchain"
	"go-blockchain/core/mempool"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"io"
	"net/http"
)

type ApiController interface {
	GetChain() response.ChainResponse
	AddTransaction(request *http.Request) interface{}
}

type ApiControllerImpl struct {
}

func (cr *ApiControllerImpl) GetChain() response.ChainResponse {
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

func (cr *ApiControllerImpl) AddTransaction(r *http.Request) interface{} {
	reqBody, _ := io.ReadAll(r.Body)
	app.Logger.Info.Log("Add Transaction Request Body: ", string(reqBody))
	txnRequest := request.TransactionRequest{}
	if err := json.Unmarshal(reqBody, &txnRequest); err != nil {
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	txn := transaction.NewTransaction(txnRequest)
	txn.SetMiningStatus(domain.READY_FOR_MINING)
	mempool.Mempool.Save(txn)
	return response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
	}
}
