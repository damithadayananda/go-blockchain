package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	coreBlock "go-blockchain/core/block"
	"go-blockchain/core/blockchain"
	"net/http/httptest"
	"testing"
)

type MockBlockChain struct {
	AddBlockFunc func(block []coreBlock.Block) error
}

func (m MockBlockChain) AddBlock(block []coreBlock.Block) error {
	return m.AddBlockFunc(block)
}

func (m MockBlockChain) ValidateBlock(block *coreBlock.Block) error {
	//TODO implement me
	panic("implement me")
}

func (m MockBlockChain) GetChain() ([]coreBlock.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBlockChain) GetLastBlock() (coreBlock.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBlockChain) DistributeBlock(block *coreBlock.Block) error {
	//TODO implement me
	panic("implement me")
}

func (m MockBlockChain) GetBlocks(noOfBlocks int) ([]coreBlock.Block, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockBlockChain) SyncBlocks(blocks []coreBlock.Block) error {
	return nil
}

// Define a mock PreProcessor struct
type MockPreProcessor struct {
	ProcessBlockFunc func(block *coreBlock.Block, callerAddress string) (error, []coreBlock.Block)
}

// Implement the IPreProcessor interface
func (m *MockPreProcessor) ProcessBlock(block *coreBlock.Block, callerAddress string) (error, []coreBlock.Block) {
	return m.ProcessBlockFunc(block, callerAddress)
}

func TestBlockControllerImpl_AddBlock_Success(t *testing.T) {
	// Create a mock pre-processor with desired behavior
	mockPreProcessor := &MockPreProcessor{
		ProcessBlockFunc: func(block *coreBlock.Block, callerAddress string) (error, []coreBlock.Block) {
			return nil, []coreBlock.Block{} // Simulate successful block processing
		},
	}

	// Create the controller and pass the mock pre-processor
	blockController := BlockControllerImpl{PreProcessor: mockPreProcessor}

	// Mock a valid block request
	blockRequest := request.BlockRequest{
		Block: request.Block{},
		Metadata: map[string]interface{}{
			"caller_address": "",
		},
	}
	reqBody, _ := json.Marshal(blockRequest)

	// Simulate HTTP request
	req := httptest.NewRequest("POST", "/add-block", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	blockchain.Chain = MockBlockChain{
		AddBlockFunc: func(block []coreBlock.Block) error {
			return nil
		},
	}
	// Call the AddBlock method
	resp := blockController.AddBlock(req)

	// Assert that the response is successful
	successResp, ok := resp.(response.SuccessResponse)
	if !ok || !successResp.Success {
		t.Errorf("expected success response, got %v", resp)
	}
}

func TestBlockControllerImpl_AddBlock_Failure(t *testing.T) {
	// Create a mock pre-processor with error behavior
	mockPreProcessor := &MockPreProcessor{
		ProcessBlockFunc: func(block *coreBlock.Block, callerAddress string) (error, []coreBlock.Block) {
			return errors.New("failed to process block"), nil // Simulate failure in block processing
		},
	}

	blockController := BlockControllerImpl{PreProcessor: mockPreProcessor}

	blockRequest := request.BlockRequest{
		Block: request.Block{},
		Metadata: map[string]interface{}{
			"caller_address": "",
		},
	}
	reqBody, _ := json.Marshal(blockRequest)

	// Simulate HTTP request
	req := httptest.NewRequest("POST", "/add-block", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	//
	blockchain.Chain = MockBlockChain{
		AddBlockFunc: func(block []coreBlock.Block) error {
			return nil
		},
	}
	// Call the AddBlock method
	resp := blockController.AddBlock(req)

	// Assert that the response is a failure
	failResp, ok := resp.(response.FailResponse)
	if !ok || failResp.Success {
		t.Errorf("expected failure response, got %v", resp)
	}
}
