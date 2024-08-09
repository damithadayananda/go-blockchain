package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/controller/request"
	"go-blockchain/controller/response"
	"go-blockchain/core/node"
	"io"
	"net/http"
	"time"
)

type NodeController interface {
	AddNode(r *http.Request) interface{}
	GetNode(r *http.Request) interface{}
}

type NodeControllerImpl struct {
}

func (cr *NodeControllerImpl) AddNode(r *http.Request) interface{} {
	reqBody, _ := io.ReadAll(r.Body)
	app.Logger.Info.Log("AddNode Request", string(reqBody))
	addNodeReq := request.AddNodeRequest{}
	if err := json.Unmarshal(reqBody, &addNodeReq); err != nil {
		app.Logger.Error.Log("Unmarshal error", err)
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	err := node.NodeRef.SaveNode(addNodeReq.Url)
	if err != nil {
		app.Logger.Error.Log("Error saving addNode", err)
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	// once node is saved
	// then need to distribute to other nodes in the network
	go cr.distributingNodeDetails(addNodeReq)
	return response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			Success: true,
		},
	}
}

func (cr *NodeControllerImpl) distributingNodeDetails(nodeRequest request.AddNodeRequest) {
	knownNodes, err := node.NodeRef.GetNodes()
	if err != nil {
		return
	}
	nodesToBeInformed := getNodesToBeInformed(nodeRequest.InformedNodes, knownNodes)
	for _, url := range nodesToBeInformed {
		body, _ := json.Marshal(request.AddNodeRequest{
			Url:           nodeRequest.Url,
			InformedNodes: knownNodes,
		})
		req, _ := http.NewRequest(http.MethodPost, url+"/node/add", bytes.NewReader(body))
		client := http.Client{
			Timeout: time.Duration(time.Second * time.Duration(config.AppConfig.NodeDistributionTimeOut)),
		}
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Request to: %s, Response: %v, Error: %v", url, string(data), err))
	}
}

func getNodesToBeInformed(receivedNodes, savedNodes []string) []string {
	nodes := make([]string, 0)
	for _, v := range savedNodes {
		for _, node := range receivedNodes {
			if v == node {
				continue
			}
		}
		if v == fmt.Sprintf("%v:%v", config.AppConfig.Host, config.AppConfig.Port) {
			continue
		}
		nodes = append(nodes, v)
	}
	return nodes
}

func (cr *NodeControllerImpl) GetNode(r *http.Request) interface{} {
	nodes, err := node.NodeRef.GetNodes()
	if err != nil {
		app.Logger.Error.Log("GetNode error", err)
		return response.FailResponse{
			BaseResponse: response.BaseResponse{
				Success: false,
			},
			Error: err.Error(),
		}
	}
	return response.NodeResponse{
		SuccessResponse: response.SuccessResponse{
			BaseResponse: response.BaseResponse{
				Success: true,
			},
		},
		Result: nodes,
	}
}
