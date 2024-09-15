package controller

import (
	"encoding/json"
	"go-blockchain/app"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/mempool"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"io"
	"net/http"
)

type TransactionController interface {
	AddTransaction(r *http.Request) interface{}
}

type TransactionControllerImpl struct {
}

func (cr *TransactionControllerImpl) AddTransaction(r *http.Request) interface{} {
	reqBody, _ := io.ReadAll(r.Body)
	app.Logger.Info.Log("Add Transaction Request Body: ", string(reqBody))
	txnRequest := request.TransactionRequest{}
	if err := json.Unmarshal(reqBody, &txnRequest); err != nil {
		app.Logger.Error.Log("Unmarshal err: ", err)
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	txn := transaction.NewTransaction(transaction.Transaction{
		Amount:   txnRequest.Amount,
		Receiver: txnRequest.Receiver,
		Sender:   txnRequest.Sender,
		Fee:      txnRequest.Fee,
	})
	txn.SetMiningStatus(domain.READY_FOR_MINING)
	mempool.Mempool.Save(txn)
	return response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
	}
}
