package main

import (
	"go-blockchain/core/blockchain"
	"go-blockchain/core/mempool"
	"go-blockchain/interface/http"
	"go-blockchain/interface/inmemorydb"
)

func main() {
	//initializing chain db
	chainDb := inmemorydb.NewInMemoryChain()
	//initializing blockchain
	blockchain.NewBlockChain(chainDb)
	//initializing transaction dn
	transactionDb := inmemorydb.NewInMemoryMemPool()
	//initializing mem pool
	mempool.NewMemPool(transactionDb)
	//initializing http server
	http.InitServer()
}
