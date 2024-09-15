package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-blockchain/app"
	pre_processor "go-blockchain/controller/pre-processor"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/blockchain"
	"io"
	"net/http"
	"strconv"
)

type BlockController interface {
	AddBlock(r *http.Request) interface{}
	GetBlocks(r *http.Request) interface{}
}

type BlockControllerImpl struct {
	PreProcessor pre_processor.IPreProcessor
}

func (bc BlockControllerImpl) AddBlock(r *http.Request) interface{} {
	reqBody, _ := io.ReadAll(r.Body)
	addBlockReq := request.BlockRequest{}
	if err := json.Unmarshal(reqBody, &addBlockReq); err != nil {
		app.Logger.Error.Log(fmt.Sprintf("Add Block Request: %v Unmarshal error: %v", string(reqBody), err))
		return handleError(err.Error())
	}
	app.Logger.Info.Log(fmt.Sprintf("Add Block Request: %v", addBlockReq))
	//validate block payload
	validationErr := addBlockReq.Validate()
	if validationErr != nil {
		return handleError(validationErr.Error())
	}
	coreBlock, _ := addBlockReq.ToCoreBlock()
	//pre-processor
	callerAddress, ok := addBlockReq.Metadata["caller_address"].(string)
	if !ok {
		return handleError("caller address is not a string")
	}
	preProcessError, blocks := bc.PreProcessor.ProcessBlock(&coreBlock, callerAddress)
	if preProcessError != nil {
		return handleError(preProcessError.Error())
	}
	err := blockchain.Chain.AddBlock(blocks)
	if err != nil {
		return handleError(err.Error())
	}
	return response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
	}
}

func handleError(str string) response.FailResponse {
	return response.FailResponse{
		BaseResponse: response.BaseResponse{
			Success: false,
		},
		Error: str,
	}
}

func (bc BlockControllerImpl) GetBlocks(r *http.Request) interface{} {
	vars := mux.Vars(r)
	noOfBlocks, err := strconv.Atoi(vars["no-of-blocks"])
	if err != nil {
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	blocks, err := blockchain.Chain.GetBlocks(noOfBlocks)
	if err != nil {
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
		}
	}
	return response.BlockResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
		Result: response.ToResponseBlocks(blocks),
	}
}
