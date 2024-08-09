package main

import (
	"go-blockchain/config"
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
	//initializing chain db
	chainDb := inmemorydb.NewInMemoryChain()
	//initializing blockchain
	blockchain.NewBlockChain(chainDb)
	//initializing transaction dn
	transactionDb := inmemorydb.NewInMemoryMemPool()
	//initializing mem pool
	mempool.NewMemPool(transactionDb)
	// starting minor
	minor.NewMinor()
	// initializing node db
	nodeDb := inmemorydb.NewInMemoryNode()
	//initializing node
	node.NewNode(nodeDb)
	//initializing http server
	http.InitServer()
}
