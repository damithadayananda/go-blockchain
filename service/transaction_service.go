package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/controller/request"
	"go-blockchain/core/blockchain"
	"go-blockchain/core/mempool"
	"go-blockchain/core/node"
	"go-blockchain/core/persistant"
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
	"go-blockchain/util"
	"io"
	"net/http"
	"time"
)

type TransactionService struct {
	MemPool mempool.MemPoolInterface
	NodeDb  persistant.NodeDBInterface
}

func NewTransactionService(mPool mempool.MemPoolInterface, nodeDb persistant.NodeDBInterface) TransactionService {
	return TransactionService{
		MemPool: mPool,
		NodeDb:  nodeDb,
	}
}

func (ts *TransactionService) AddTransaction(txn transaction.Transaction) {
	doubleSpending := false
	// if the given txn is in mined block will ignore it
	blocks, _ := blockchain.Chain.GetChain()
	for _, block := range blocks {
		if txns, ok := block.Data.([]transaction.Transaction); ok {
			for _, t := range txns {
				if t.Id == txn.Id {
					doubleSpending = true
					break
				}
			}
			if doubleSpending {
				break
			}
		}
	}
	// if the txn is in mined block will ignore it
	memPoolTxns, _ := ts.MemPool.GetAll()
	if len(memPoolTxns) > 0 {
		for _, v := range memPoolTxns {
			if v.Id == txn.Id {
				doubleSpending = true
			}
		}
	}
	// not double spending means new txn
	//if double spending no action taken
	if !doubleSpending {
		//save transaction to mempool
		ts.MemPool.Save(txn)
		//distribute txn
		//distributing txn prior to updating node is important
		//otherwise new node txn also will be delivered to new node
		distributeTransaction(txn)
		//for node validator txn
		if len(txn.Data) > 0 {
			address, _ := txn.Data["Address"].(string)
			status, _ := txn.Data["Status"].(string)
			st := domain.NodeStatusFromString(status)
			ts.NodeDb.UpdateNodeStatus(address, st)
		}
	}
}

func distributeTransaction(txn transaction.Transaction) {
	knownNodes, err := node.NodeRef.GetNodes()
	if err != nil {
		return
	}
	body, _ := json.Marshal(request.TransactionRequest{
		Amount:   txn.Amount,
		Receiver: txn.Receiver,
		Sender:   txn.Sender,
		Fee:      txn.Fee,
		Id:       txn.Id,
		Data:     txn.Data,
	})
	for _, node := range knownNodes {
		if node.Status != domain.ACTIVE {
			// new transactions will be distributed among only active nodes
			app.Logger.Info.Log(fmt.Sprintf("Node %s is not active, skipping...", node.Ip))
			continue
		}
		req, _ := http.NewRequest(http.MethodPost, node.Ip+"/transaction/add", bytes.NewReader(body))
		client := util.GeHttpsClient(node.Certificate)
		client.Timeout = time.Second * time.Duration(config.AppConfig.NodeDistributionTimeOut)
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Request to: %s, Response: %v, Error: %v", node.Ip, string(data), err))
	}
}
