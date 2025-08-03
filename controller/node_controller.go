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
	"go-blockchain/domain"
	"go-blockchain/util"
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
	err := node.NodeRef.SaveNode(domain.Node{
		Ip:          addNodeReq.Url,
		Certificate: addNodeReq.Certificate,
		Address:     addNodeReq.Address,
	})
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
	iPsToBeInformed := getNodesToBeInformed(extractNodeIPs(nodeRequest.InformedNodes), extractNodeIPs(knownNodes))
	for _, node := range getNodeFromIp(knownNodes, iPsToBeInformed) {
		body, _ := json.Marshal(request.AddNodeRequest{
			Url:           node.Ip,
			InformedNodes: append(nodeRequest.InformedNodes, createNodesToBeInformed(knownNodes, iPsToBeInformed)...),
			Certificate:   node.Certificate,
			Address:       node.Address,
		})
		req, _ := http.NewRequest(http.MethodPost, node.Ip+"/node/add", bytes.NewReader(body))
		client := util.GeHttpsClient(node.Certificate)
		client.Timeout = time.Second * time.Duration(config.AppConfig.NodeDistributionTimeOut)
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Request to: %s, Response: %v, Error: %v", node.Ip, string(data), err))
	}
}

func getNodeFromIp(nodes []domain.Node, urls []string) []domain.Node {
	var nodeList []domain.Node
	for _, ip := range urls {
		for _, node := range nodes {
			if ip == node.Ip {
				nodeList = append(nodeList, node)
			}
		}
	}
	return nodeList
}

func createNodesToBeInformed(knownNodes []domain.Node, ips []string) []domain.Node {
	var nodes []domain.Node
	for _, v := range ips {
		for _, k := range knownNodes {
			if v == k.Ip {
				nodes = append(nodes, k)
			}
		}
	}
	return nodes
}

func extractNodeIPs(nodes []domain.Node) []string {
	ips := make([]string, len(nodes))
	for _, v := range nodes {
		ips = append(ips, v.Ip)
	}
	return ips
}

func getNodesToBeInformed(receivedNodes, savedNodes []string) []string {
	nodes := make([]string, 0)
	for _, v := range savedNodes {
		var found bool
		for _, node := range receivedNodes {
			if v == node {
				found = true
			}
		}
		//if v == fmt.Sprintf("%v:%v", config.AppConfig.Host, config.AppConfig.Port) {
		//	continue
		//}
		if !found {
			nodes = append(nodes, v)
		}
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
		Result: fromDomainNodeToNode(nodes),
	}
}

func fromDomainNodeToNode(nodes []domain.Node) []response.Node {
	var res []response.Node
	for _, v := range nodes {
		res = append(res, response.Node{
			Ip:          v.Ip,
			Certificate: v.Certificate,
			Address:     v.Address,
		})
	}
	return res
}
