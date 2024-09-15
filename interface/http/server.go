package http

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-blockchain/config"
	"go-blockchain/controller"
	pre_processor "go-blockchain/controller/pre-processor"
	"go-blockchain/interface/http/api"
	"go-blockchain/interface/http/ui"
	"log"
	"net/http"
)

func InitServer() {
	mux := mux.NewRouter()
	apiHandler := api.API{}
	apiHandler.RegisterAPI(mux,
		&controller.TransactionControllerImpl{},
		&controller.ChainControllerImpl{},
		&controller.NodeControllerImpl{},
		&controller.BlockControllerImpl{
			PreProcessor: pre_processor.IndexValidator{},
		})
	uiHandler := ui.UI{}
	uiHandler.RegisterUi(mux)

	// Configure CORS
	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handler := handlers.CORS(corsOptions, corsMethods, corsHeaders)(mux)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.Port), handler); err != nil {
		log.Fatalf(err.Error())
	}
}
