package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-blockchain/app"
	"go-blockchain/config"
	"go-blockchain/controller/request"
	"go-blockchain/core/persistant"
	"io"
	"net/http"
	"strconv"
	"time"
)

var NodeRef *Node

type Node struct {
	database persistant.NodeDBInterface
}

func NewNode(database persistant.NodeDBInterface) {
	NodeRef = &Node{
		database: database,
	}
	// Adding known nodes from config
	for _, node := range config.AppConfig.KnownNodes {
		database.Save(node)
	}
	// distribute details to the cluster
	NodeRef.distributeNodeExistence()
}

func (n *Node) SaveNode(url string) error {
	//avoid duplicate node details
	data, err := n.database.GetAll()
	if err != nil {
		return err
	}
	for _, v := range data {
		if v == url {
			return nil
		}
	}
	//saving new node
	return n.database.Save(url)
}

func (n *Node) GetNodes() ([]string, error) {
	return n.database.GetAll()
}

func (n *Node) RemoveNode(url string) error {
	return n.database.Delete(url)
}

// implementing gossip protocol
// once node is initiated it will inform it's existence to known nodes
// these known nodes will inform their known nodes about newly join node
func (n *Node) distributeNodeExistence() {
	myHost := config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port)
	informedNodes, err := n.GetNodes()
	if err != nil {
		return
	}
	informedNodes = append(informedNodes, myHost)
	for _, node := range config.AppConfig.KnownNodes {
		body, _ := json.Marshal(request.AddNodeRequest{
			Url:           config.AppConfig.Host + ":" + strconv.Itoa(config.AppConfig.Port),
			InformedNodes: informedNodes,
		})
		req, _ := http.NewRequest(http.MethodPost, node+"/node/add", bytes.NewReader(body))
		client := http.Client{
			Timeout: time.Second * time.Duration(config.AppConfig.NodeDistributionTimeOut),
		}
		res, err := client.Do(req)
		var data []byte
		if res != nil {
			data, _ = io.ReadAll(res.Body)
			res.Body.Close()
		}
		app.Logger.Info.Log(fmt.Sprintf("Request to: %s, Response: %v, Error: %v", node, string(data), err))
	}
}
