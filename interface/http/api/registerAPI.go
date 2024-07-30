package api

import (
	"encoding/json"
	"go-blockchain/controller"
	"net/http"
)

type API struct {
	controller controller.ApiController
}

func (api *API) RegisterAPI(mux *http.ServeMux, controller controller.ApiController) {
	api.controller = controller
	mux.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(api.controller.GetChain())
		w.Write(data)
	})
	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {

	})
}
