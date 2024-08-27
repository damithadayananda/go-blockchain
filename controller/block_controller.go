package controller

import (
	"encoding/json"
	"go-blockchain/app"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/blockchain"
	"io"
	"net/http"
)

type BlockController interface {
	AddBlock(r *http.Request) interface{}
}

type BlockControllerImpl struct {
}

func (bc BlockControllerImpl) AddBlock(r *http.Request) interface{} {
	reqBody, _ := io.ReadAll(r.Body)
	addBlockReq := request.Block{}
	if err := json.Unmarshal(reqBody, &addBlockReq); err != nil {
		app.Logger.Error.Log("Unmarshal error", err)
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	coreBlock, _ := addBlockReq.ToCoreBlock()
	err := blockchain.Chain.AddBlock(&coreBlock)
	if err != nil {
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	return response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
	}
}
