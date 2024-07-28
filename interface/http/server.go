package http

import (
	"go-blockchain/interface/http/ui"
	"log"
	"net/http"
)

func InitServer() {
	mux := http.NewServeMux()
	ui.RegisterUi(mux)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf(err.Error())
	}
}
