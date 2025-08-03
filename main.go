package main

import (
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/controller"
	"go-blockchain/core/blockchain"
	"go-blockchain/core/mempool"
	"go-blockchain/core/minor"
	"go-blockchain/core/node"
	"go-blockchain/interface/http"
	"go-blockchain/interface/inmemorydb"
)

func main() {
	//initializing configuration
	config.InitConfig()
	//initializing application
	app.NewApplication()
	//initializing chain db
	chainDb := inmemorydb.NewInMemoryChain()
	//initializing blockchain
	blockchain.NewBlockChain(chainDb)
	//initializing transaction dn
	transactionDb := inmemorydb.NewInMemoryMemPool()
	//initializing mem pool
	memPoolRef := mempool.NewMemPool(transactionDb)
	// starting minor
	minor.NewMinor(memPoolRef)
	// initializing node db
	nodeDb := inmemorydb.NewInMemoryNode()
	//initializing node
	node.NewNode(nodeDb)
	//controllers
	txnCtrl := controller.NewTransactionController(memPoolRef)
	//initializing http server
	http.InitServer(txnCtrl)
}
