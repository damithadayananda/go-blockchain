package pre_processor

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-blockchain/config"
	"go-blockchain/controller/response"
	coreBlock "go-blockchain/core/block"
	"go-blockchain/core/blockchain"
	"io"
	"net/http"
	"time"
)

type IndexValidator struct {
}

func (iv IndexValidator) ProcessBlock(block *coreBlock.Block, callerAddress string) (error, []coreBlock.Block) {
	latestBlock, err := blockchain.Chain.GetLastBlock()
	if err != nil {
		return err, nil
	}
	if block.Index == latestBlock.Index {
		// This is a conflict
		// for conflicts, blocks with most transaction will be remained in chain
		existingDataLength := getLengthOfData(latestBlock.Data)
		newDataLength := getLengthOfData(block.Data)
		if existingDataLength < newDataLength {
			return nil, []coreBlock.Block{*block}
		}
		return errors.New("existing or outdated block"), nil
	} else if block.Index > latestBlock.Index+1 {
		// our chain is behind, need to ask for missing blocks
		// here idea is calling back to the caller and ask for missing blocks
		// avoid to distributing fraudulent transactions transaction validations need to be placed
		httpClient := http.Client{
			Timeout: time.Duration(config.AppConfig.BlockInquiryTimeOut) * time.Second,
		}
		missingNoOfBlocks := block.Index - latestBlock.Index
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/block/%d", callerAddress, missingNoOfBlocks), nil)
		if err != nil {
			return err, nil
		}
		resp, err := httpClient.Do(request)
		if err != nil {
			return err, nil
		}
		var result []response.Block
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err, nil
		}
		err = json.Unmarshal(data, &result)
		if err != nil {
			return err, nil
		}
		blocks := response.ToCoreBlocks(result)
		return nil, blocks
	} else if block.Index == latestBlock.Index+1 {
		return nil, []coreBlock.Block{*block}
	}
	return errors.New("block is already added or invalid block"), nil
}

func getLengthOfData(data interface{}) int {
	switch v := data.(type) {
	case []interface{}:
		return len(v)
	case string:
		return len(v)
	}
	return 0
}
