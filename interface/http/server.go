package http

import (
	"go-blockchain/controller"
	"go-blockchain/interface/http/api"
	"go-blockchain/interface/http/ui"
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()
	apiHandler := api.API{}
	apiHandler.RegisterAPI(mux, &controller.ApiControllerImpl{})
	uiHandler := ui.UI{}
	uiHandler.RegisterUi(mux)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf(err.Error())
	}
}
