package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-blockchain/controller"
	"net/http"
)

type API struct {
	transactionController controller.TransactionController
	chainController       controller.ChainController
	nodeController        controller.NodeController
	blockController       controller.BlockController
}

func (api *API) RegisterAPI(mux *mux.Router,
	trCtr controller.TransactionController,
	cCtr controller.ChainController,
	nCtr controller.NodeController,
	bCtr controller.BlockController) {
	api.transactionController = trCtr
	api.chainController = cCtr
	api.nodeController = nCtr
	api.blockController = bCtr

	mux.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.chainController.GetChain())
		w.Write(data)
	})
	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.transactionController.AddTransaction(r))
		w.Write(data)
	})

	// node routes
	node := mux.PathPrefix("/node").Subrouter()
	node.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.nodeController.AddNode(r))
		w.Write(data)
	})
	node.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.nodeController.GetNode(r))
		w.Write(data)
	})

	// block routes
	block := mux.PathPrefix("/block").Subrouter()
	block.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.blockController.AddBlock(r))
		w.Write(data)
	})
	block.HandleFunc("/{no-of-blocks}", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.blockController.GetBlocks(r))
		w.Write(data)
	})
}
