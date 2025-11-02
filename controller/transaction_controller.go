package controller

import (
	"encoding/json"
	"go-blockchain/app"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/mempool"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"go-blockchain/service"
	"io"
	"net/http"
)

type TransactionController interface {
	AddTransaction(r *http.Request) interface{}
	GetMemPoolTransactions(r *http.Request) interface{}
}

type TransactionControllerImpl struct {
	MemPool mempool.MemPoolInterface
	TxnSvc  service.TransactionService
}

func NewTransactionController(mPool mempool.MemPoolInterface, txnSvc service.TransactionService) TransactionController {
	return &TransactionControllerImpl{
		MemPool: mPool,
		TxnSvc:  txnSvc,
	}
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
		Id:       txnRequest.Id,
		Data:     txnRequest.Data,
	})
	txn.SetMiningStatus(domain.READY_FOR_MINING)
	cr.TxnSvc.AddTransaction(txn)
	return response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
	}
}

func (cr *TransactionControllerImpl) GetMemPoolTransactions(r *http.Request) interface{} {
	txns, err := cr.MemPool.GetAll()
	if err != nil {
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	responseTransaction := []response.Transaction{}
	for _, txn := range txns {
		t := response.Transaction{}
		t.FromDomain(txn)
		responseTransaction = append(responseTransaction, t)
	}
	return response.TransactionResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
		Result: responseTransaction,
	}
}
