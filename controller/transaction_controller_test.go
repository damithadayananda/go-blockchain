package controller

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"net/http/httptest"
	"testing"
)

type MockMemPool struct {
	SaveMock func(transaction transaction.Transaction) error
}

func (m MockMemPool) Save(transaction transaction.Transaction) error {
	return m.SaveMock(transaction)
}

func (m MockMemPool) Get() ([]transaction.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockMemPool) Delete(ids []string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockMemPool) Mark(transactions []transaction.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func TestAddTransactionSuccess(t *testing.T) {
	txnCtrl := NewTransactionController(MockMemPool{
		SaveMock: func(transaction transaction.Transaction) error {
			return nil
		},
	})
	txnReq := request.TransactionRequest{
		Amount:       20,
		Receiver:     "",
		Sender:       "",
		Fee:          10,
		MiningStatus: domain.READY_FOR_MINING,
	}
	reqBody, _ := json.Marshal(txnReq)
	req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := txnCtrl.AddTransaction(req)
	respBody, _ := resp.(response.SuccessResponse)
	assert.Equal(t, true, respBody.Success)
}

func TestAddTransactionFailure(t *testing.T) {
	txnCtrl := NewTransactionController(MockMemPool{
		SaveMock: func(transaction transaction.Transaction) error {
			return nil
		},
	})
	txnReq := ""
	reqBody, _ := json.Marshal(txnReq)
	req := httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	resp := txnCtrl.AddTransaction(req)
	respBody, _ := resp.(response.FailResponse)
	assert.Equal(t, false, respBody.Success)
}
