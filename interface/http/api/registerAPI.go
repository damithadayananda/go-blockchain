package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-blockchain/controller"
	"net/http"
)

type API struct {
	controller controller.ApiController
}

func (api *API) RegisterAPI(mux *mux.Router, controller controller.ApiController) {
	api.controller = controller
	mux.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.controller.GetChain())
		w.Write(data)
	})
	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.controller.AddTransaction(r))
		w.Write(data)
	})
}
